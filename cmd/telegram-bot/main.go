package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tbot "github.com/ma11oc/go-predictor/internal/tbot"
	pb "github.com/ma11oc/go-predictor/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Request is used to store users queries
type Request struct {
	gorm.Model
	ChatID           int64
	MessageID        int
	UserID           int
	Command          string
	CommandArguments string
	Response         *Response
}

// Response is used to store response payload
type Response struct {
	gorm.Model
	RequestID int
	MessageID int
	ChatID    int64
	TTL       int
	Payload   []byte
}

// User is used to store users info
type User struct {
	gorm.Model
	ID           int `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	UserName     string
	LanguageCode string
	IsBot        bool
}

func main() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("BOT_DB_PATH")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Request{}, &Response{}, &User{})

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// predictor service client

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	var srvAddr string
	var ok bool

	if srvAddr, ok = os.LookupEnv("PREDICTOR_SERVER_ADDRESS"); !ok {
		srvAddr = "localhost:50051"
	}

	conn, err := grpc.Dial(srvAddr, opts...) // FIXME: hardcoded address
	if err != nil {
		log.Panicf("unable to communicate with predictor server: %s", err)
	}
	defer conn.Close()

	psc := pb.NewPredictorServiceClient(conn)

	// main loop
	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil && update.InlineQuery == nil { // ignore any non-Message Updates
			continue
		}

		if update.CallbackQuery != nil {
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))

			row := &Response{}

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
		}

		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			// var err error

			if update.Message.IsCommand() {
				log.Printf("got command: %s", update.Message.Command())

				switch update.Message.Command() {
				case "new":
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

					resp, err := psc.ComputePerson(context.Background(), req)
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

					user := &User{
						ID:           update.Message.From.ID,
						FirstName:    update.Message.From.FirstName,
						LastName:     update.Message.From.LastName,
						UserName:     update.Message.From.UserName,
						LanguageCode: update.Message.From.LanguageCode,
						IsBot:        update.Message.From.IsBot,
					}

					db.FirstOrCreate(user)
					db.Create(&Request{
						ChatID:           update.Message.Chat.ID,
						MessageID:        update.Message.MessageID,
						UserID:           user.ID,
						Command:          update.Message.Command(),
						CommandArguments: update.Message.CommandArguments(),
						Response: &Response{
							TTL:       30,
							MessageID: feedback.MessageID,
							ChatID:    feedback.Chat.ID,
							Payload:   personBytes,
						},
					})

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
