package cmd

import (
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ma11oc/go-predictor/pkg/tracer"

	"context"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tbot "github.com/ma11oc/go-predictor/internal/tbot"
	model "github.com/ma11oc/go-predictor/internal/tbot/db"
	pb "github.com/ma11oc/go-predictor/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbCfg Database
var pSrvCfg PredictorServer
var authCfg Auth

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("database.connection_string: %s\n", viper.GetString("database.connection_string"))

		// TODO: validate tracer config

		sn := viper.GetString("tracer.service_name")
		ce := viper.GetString("tracer.endpoint")

		tracer, closer, err := tracer.NewGlobalTracer(sn, ce)
		if err != nil {
			panic(err)
		}
		defer closer.Close()

		// tracer := opentracing.GlobalTracer()

		// span := tracer.StartSpan("say-hello")
		// println("hello")
		// span.Finish()

		db, err := gorm.Open(sqlite.Open(viper.GetString("database.connection_string")), &gorm.Config{})
		if err != nil {
			log.Fatalf("database path: %s", viper.GetString("database.connection_string"))
			log.Panic(err)
		}

		// Migrate the schema
		db.AutoMigrate(&model.Request{}, &model.Response{}, &model.User{})

		bot, err := tgbotapi.NewBotAPI(viper.GetString("auth.token"))
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		// predictor service client

		unaryInterceptor := grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(),
		)

		opts := []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(unaryInterceptor),
		}

		conn, err := grpc.Dial(viper.GetString("predictor.endpoint"), opts...) // FIXME: hardcoded address
		if err != nil {
			log.Panicf("unable to communicate with predictor server: %s", err)
		}
		defer conn.Close()

		psc := pb.NewPredictorServiceClient(conn)

		runPredictorBot(db, bot, psc, tracer)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().BoolVar(&dbCfg.AutoMigrate, "db-auto-migrate", true, "apply migrations (default is `true`)")
	viper.BindPFlag("database.auto_migrate", runCmd.PersistentFlags().Lookup("db-auto-migrate"))

	runCmd.PersistentFlags().StringVar(&dbCfg.ConnectionString, "db-conn-str", fmt.Sprintf("sqlite://%s.db", ProgramName), "db connection string")
	viper.BindPFlag("database.connection_string", runCmd.PersistentFlags().Lookup("db-conn-str"))

	runCmd.PersistentFlags().StringVar(&pSrvCfg.Endpoint, "predictor-srv-url", "", "url of predictor server (default is localhost:50051)")
	viper.BindPFlag("predictor.endpoint", runCmd.PersistentFlags().Lookup("predictor-srv-url"))

	runCmd.PersistentFlags().StringVar(&authCfg.Token, "auth-token", "", "telegram auth token")
	viper.BindPFlag("auth.token", runCmd.PersistentFlags().Lookup("auth-token"))

	// fmt.Println("run init done")
}

func runPredictorBot(db *gorm.DB, bot *tgbotapi.BotAPI, psc pb.PredictorServiceClient, tracer opentracing.Tracer) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// main loop
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil && update.InlineQuery == nil { // ignore any non-Message Updates
			continue
		}

		if update.CallbackQuery != nil {
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))

			row := &model.Response{}

			db.Where("chat_id = ? AND message_id = ?", update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID).First(row)

			person := &pb.Person{}

			if err := proto.Unmarshal(row.Payload, person); err != nil {
				log.Fatalf("unable to unmarshal result: %s", err)
			}

			msg := tgbotapi.NewEditMessageText(
				update.CallbackQuery.Message.Chat.ID,
				update.CallbackQuery.Message.MessageID,
				"",
			)
			msg.ParseMode = tgbotapi.ModeHTML

			msg.Text, err = tbot.MakeMessageByCallback(person, update.CallbackQuery.Data)
			if err != nil {
				msg.Text = fmt.Sprintf("Error: %s", err)
				break
			}

			markup, err := tbot.MakePersonMarkup(person)
			if err != nil {
				msg.Text = fmt.Sprintf("Error: %s", err)
				break
			}

			msg.ReplyMarkup = &markup

			if msg.Text != "" {
				bot.Send(msg)
			}
			// bot.Send(
			// tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, update.CallbackQuery.Message.ReplyMarkup))

		}

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.InlineQuery != nil {
			switch update.InlineQuery.Query {
			case "card":
				log.Printf("inline query: %s", update.InlineQuery.Query)
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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			// var err error

			if update.Message.IsCommand() {
				log.Printf("got command: %s", update.Message.Command())

				switch update.Message.Command() {
				case "new":

					span := tracer.StartSpan("command.new")

					log.Printf("command arguments: %s", update.Message.CommandArguments())

					args := strings.Split(update.Message.CommandArguments(), " ")

					_, err := time.Parse("2006-01-02", args[0])
					if err != nil {
						msg.Text = fmt.Sprintf("Error: %s", err)
						bot.Send(msg)
						break
					}

					year, err := strconv.Atoi(args[1])
					if err != nil {
						msg.Text = fmt.Sprintf("Error: %s", err)
						bot.Send(msg)
						break
					}

					req := &pb.PersonRequest{
						Api:  "v1",
						Lang: "ru-RU",

						PersonProfile: &pb.PersonProfile{
							Name:     "A Person",
							Gender:   1,
							Features: 0,
							Birthday: args[0],
							Age:      int32(year),
						},
					}

					ctx, cancel := context.WithTimeout(opentracing.ContextWithSpan(context.Background(), span), time.Second*3)
					defer cancel()

					resp, err := psc.ComputePerson(ctx, req)
					if err != nil {
						log.Fatalln("got error during communication to the server: ", err)
					}

					person := resp.GetPerson()
					personBytes, err := proto.Marshal(person)
					if err != nil {
						log.Fatalf("unable to marshal person")
					}

					msg.Text, _ = tbot.MakeMessageByCallback(person, "card:base:main:desc") // FIXME: error handling
					markup, _ := tbot.MakePersonMarkup(person)                              // FIXME: error handling
					msg.ReplyMarkup = markup
					msg.ParseMode = tgbotapi.ModeHTML

					feedback, _ := bot.Send(msg)

					/*
					 * user := &User{
					 *     User: update.Message.From,
					 * }
					 */

					/*
					 * db.Create(&Request{
					 *     ChatID:    update.Message.Chat.ID,
					 *     MessageID: feedback.MessageID,
					 *     Result:    personBytes,
					 *     UserID:    user.User.ID,
					 * })
					 */

					// Create a Child Span. Note that we're using the ChildOf option.
					childSpan := tracer.StartSpan(
						"db.create",
						opentracing.ChildOf(span.Context()),
					)

					user := &model.User{
						ID:           update.Message.From.ID,
						FirstName:    update.Message.From.FirstName,
						LastName:     update.Message.From.LastName,
						UserName:     update.Message.From.UserName,
						LanguageCode: update.Message.From.LanguageCode,
						IsBot:        update.Message.From.IsBot,
					}

					// TODO: transaction
					db.FirstOrCreate(user)
					db.Create(&model.Request{
						ChatID:           update.Message.Chat.ID,
						MessageID:        update.Message.MessageID,
						UserID:           user.ID,
						Command:          update.Message.Command(),
						CommandArguments: update.Message.CommandArguments(),
						Response: &model.Response{
							TTL:       30,
							MessageID: feedback.MessageID,
							ChatID:    feedback.Chat.ID,
							Payload:   personBytes,
						},
					})

					childSpan.Finish()

					span.Finish()

				case "settings":
					log.Printf("don't know how to handle settings")

				case "help", "start":
					msg.ParseMode = tgbotapi.ModeHTML
					msg.Text = "" +
						"/new: Request for a new prediction\n" +
						"```\n" +
						"/new 1970-01-01 20 \n" +
						"```\n" +
						"/help: Show this message\n"

					bot.Send(msg)
				}
			}

		}
	}
}
