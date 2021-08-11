package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/ma11oc/go-predictor/pkg/logger"
	"github.com/ma11oc/go-predictor/pkg/tracer"

	// "log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tbot "github.com/ma11oc/go-predictor/internal/tbot"
	callbacks "github.com/ma11oc/go-predictor/internal/tbot/callbacks"
	commands "github.com/ma11oc/go-predictor/internal/tbot/commands"
	model "github.com/ma11oc/go-predictor/internal/tbot/db"
	pb "github.com/ma11oc/go-predictor/pkg/api/v1"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var cfg = NewConfig()

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		var runtimeConfig *Config
		var err error

		if runtimeConfig, err = NewConfigFromMap(viper.AllSettings()); err != nil {
			return err
		}

		if err = runtimeConfig.Validate(); err != nil {
			return err
		}

		// initialize logger
		if err := logger.Init(runtimeConfig.Logger.Level, runtimeConfig.Logger.TimeFormat); err != nil {
			return fmt.Errorf("Failed to initialize logger: %v", err)
		}
		defer logger.HandlePanic(ProgramName+".Run", logger.Log) // FIXME: bot crashes despite recovery

		// initialize tracer
		logger.Log.Info("Initializing tracer")
		tracer, closer, err := tracer.NewGlobalTracer(runtimeConfig.Tracer.ServiceName, runtimeConfig.Tracer.Endpoint)
		if err != nil {
			panic(err)
		}
		defer closer.Close()

		logger.Log.Info("Initializing database")
		db, err := gorm.Open(sqlite.Open(runtimeConfig.Database.ConnectionString), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect to db: %s", err)
		}

		// Migrate the schema
		if runtimeConfig.Database.AutoMigrate {
			logger.Log.Info("Applying database migrations")
			if err := db.AutoMigrate(&model.Request{}, &model.Response{}, &model.User{}, &model.PersonProfile{}); err != nil {
				return fmt.Errorf("Failed to apply database migrations: %s", err)
			}
		}

		logger.Log.Info("Initializing BotAPI")
		bot, err := tgbotapi.NewBotAPI(runtimeConfig.Auth.Token)
		if err != nil {
			return fmt.Errorf("Failed to initialize a new BotAPI: %s", err)
		}

		bot.Debug = true

		logger.Log.Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

		// predictor service client

		logger.Log.Info("Initializing GRPC Client")
		unaryInterceptor := grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(),
		)

		opts := []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(unaryInterceptor),
		}

		conn, err := grpc.Dial(runtimeConfig.PredictorServer.Endpoint, opts...)
		if err != nil {
			return fmt.Errorf("Failed to dial predictor server: %s", err)
		}
		defer conn.Close()

		psc := pb.NewPredictorServiceClient(conn)

		return runPredictorBot(db, bot, psc, tracer)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().BoolVar(&cfg.Database.AutoMigrate, "db-auto-migrate", true, "apply migrations")
	viper.BindPFlag("database.auto_migrate", runCmd.PersistentFlags().Lookup("db-auto-migrate"))

	runCmd.PersistentFlags().StringVar(&cfg.Database.ConnectionString, "db-conn-str", fmt.Sprintf("sqlite://%s.db", ProgramName), "db connection string")
	viper.BindPFlag("database.connection_string", runCmd.PersistentFlags().Lookup("db-conn-str"))

	runCmd.PersistentFlags().StringVar(&cfg.PredictorServer.Endpoint, "predictor-srv-url", "", "url of predictor server (default is localhost:50051)")
	viper.BindPFlag("predictor.endpoint", runCmd.PersistentFlags().Lookup("predictor-srv-url"))

	runCmd.PersistentFlags().StringVar(&cfg.Auth.Token, "auth-token", "asdf", "telegram auth token")
	viper.BindPFlag("auth.token", runCmd.PersistentFlags().Lookup("auth-token"))

	runCmd.PersistentFlags().IntVar(&cfg.Logger.Level, "log-level", 0, "log level (debug=-1, fatal=5)")
	viper.BindPFlag("logger.level", runCmd.PersistentFlags().Lookup("log-level"))

	// fmt.Println("run init done")
}

func runPredictorBot(db *gorm.DB, bot *tgbotapi.BotAPI, psc pb.PredictorServiceClient, tracer opentracing.Tracer) error {
	// TODO: move updates chan to init stage
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.Log.Fatal("Failed to get updates channel", zap.Error(err))
	}

	// main loop
	logger.Log.Info("Entering main loop")
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil && update.InlineQuery == nil { // ignore any non-Message Updates
			continue
		}

		// TODO: create context here and start span
		if update.CallbackQuery != nil {
			// TODO: test CallbackQuery.Data nil or ""
			if update.CallbackQuery.Data == "" {
				logger.Log.Warn("Received empty callback data")
				continue
			}

			callback := strings.Split(update.CallbackQuery.Data, ":")
			if len(callback) == 0 {
				logger.Log.Warn("Failed to decode callback", zap.String("callback", update.CallbackQuery.Data))
			}

			switch callback[0] {
			case "person":
				logger.Log.Debug("Received person callback")

				span := tracer.StartSpan("update.CallbackQuery")

				span.LogFields(log.String("callback", update.CallbackQuery.Data))
				span.SetTag("type", "callback")

				ctx, cancel := context.WithTimeout(opentracing.ContextWithSpan(context.Background(), span), time.Second*3)
				defer cancel()

				callbacks.HandlePersonCallbackQuery(ctx, update.CallbackQuery, psc, db, bot)

				span.Finish()
			default:
				logger.Log.Warn("Received unknown callback. Skipping...", zap.String("callback", update.CallbackQuery.Data))
				if _, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)); err != nil {
					logger.Log.Error("Failed to answer callback query", zap.Error(err))
				}
				continue
			}

		}

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.InlineQuery != nil {
			switch update.InlineQuery.Query {
			case "card":
				logger.Log.Debug("inline query", zap.String("inline query", update.InlineQuery.Query))
				resp := make([]interface{}, 5)

				for i := range []int{1, 2, 3, 4, 5} {
					resp[i] = tgbotapi.NewInlineQueryResultArticle(fmt.Sprintf("%d", i), fmt.Sprintf("%d%d", i, i), fmt.Sprintf("%d%d%d", i, i, i))
				}

				bot.AnswerInlineQuery(tgbotapi.InlineConfig{
					InlineQueryID: update.InlineQuery.ID,
					Results:       resp,
					IsPersonal:    true,
				})
			default:
				continue
			}

		}

		if update.Message != nil {
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			// var err error

			if update.Message.IsCommand() {
				logger.Log.Sugar().Debug("got command: %s", update.Message.Command())

				switch update.Message.Command() {
				case "new":
					// TODO: start main span here
					if err := commands.HandleCommandNew(update.Message, psc, db, bot); err != nil {
						logger.Log.Sugar().Error(err)

						// TODO: move into HandleCommandNew func
						errMsg, err := tbot.MakeErrorMessage("Internal server error", err)
						if err != nil {
							logger.Log.Sugar().Error(err)
						}
						errMsg.ChatID = update.Message.Chat.ID
						bot.Send(errMsg)

					}

				case "settings":
					//log.Printf("don't know how to handle settings")
					continue

				case "help", "start":
					// msg.ParseMode = tgbotapi.ModeHTML
					// msg.Text = "" +
					// 	"/new: Request for a new prediction\n" +
					// 	"```\n" +
					// 	"/new 1970-01-01 20 \n" +
					// 	"```\n" +
					// 	"/help: Show this message\n"

					// bot.Send(msg)
				}
			}

		}

	}
	return nil
}
