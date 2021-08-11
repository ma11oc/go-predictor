package tbot

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"

	"github.com/ma11oc/go-predictor/internal/tbot"
	model "github.com/ma11oc/go-predictor/internal/tbot/db"
	markups "github.com/ma11oc/go-predictor/internal/tbot/markups"
	v1 "github.com/ma11oc/go-predictor/pkg/api/v1"
	"github.com/ma11oc/go-predictor/pkg/logger"
)

// HandleCommandNew parses args and makes response
func HandleCommandNew(msg *tgbotapi.Message, psc v1.PredictorServiceClient, db *gorm.DB, bot *tgbotapi.BotAPI) error {
	var cspan opentracing.Span

	logger.Log.Sugar().Debug("command new: %s", msg.Text)

	tracer := opentracing.GlobalTracer()

	pspan := tracer.StartSpan("HandleCommandNew")
	defer pspan.Finish()

	pspan.LogFields(log.String("command new", msg.Text))
	pspan.
		SetTag("username", msg.From.UserName).
		SetTag("chat_id", msg.Chat.ID).
		SetTag("message_id", msg.MessageID)

	ctx, cancel := context.WithTimeout(opentracing.ContextWithSpan(context.Background(), pspan), time.Second*3)
	defer cancel()

	preq, err := parseRequest(msg.CommandArguments())
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	resp, err := psc.ComputePerson(ctx, preq)
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	person := resp.GetPerson()

	reply := tgbotapi.NewMessage(msg.Chat.ID, "")
	reply.Text, err = tbot.MakeMessageByCallback(person, "person:card:base:main:desc") // FIXME: MakeMessageByCallback?
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	markup, err := markups.MakePersonMarkup(person)
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	reply.ReplyMarkup = markup
	reply.ParseMode = tgbotapi.ModeHTML

	cspan = opentracing.StartSpan("bot.Send", opentracing.ChildOf(pspan.Context()))
	defer cspan.Finish()

	feedback, err := bot.Send(reply)
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	cspan = opentracing.StartSpan("db.Transaction", opentracing.ChildOf(pspan.Context()))
	defer cspan.Finish()

	user := &model.User{
		ID:           msg.From.ID,
		FirstName:    msg.From.FirstName,
		LastName:     msg.From.LastName,
		UserName:     msg.From.UserName,
		LanguageCode: msg.From.LanguageCode,
		IsBot:        msg.From.IsBot,
	}

	mreq := &model.Request{
		ChatID:           msg.Chat.ID,
		MessageID:        msg.MessageID,
		UserID:           user.ID,
		Command:          msg.Command(),
		CommandArguments: msg.CommandArguments(),
		Response: &model.Response{
			TTL:       30,
			MessageID: feedback.MessageID,
			ChatID:    feedback.Chat.ID,
			// Payload:   personBytes,
		},
		PersonProfile: &model.PersonProfile{
			PersonProfile: preq.PersonProfile,
		},
	}

	// TODO: transaction
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.FirstOrCreate(user).Error; err != nil {
			return err
		}

		if err := tx.Create(mreq).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		cspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	// span.Finish()

	return nil

}

func parseRequest(text string) (*v1.PersonRequest, error) {
	var birthday string
	var age int
	var err error

	args := strings.Split(text, " ")

	if len(args) < 2 {
		return nil, fmt.Errorf("Malformed request")
	}

	if len(args) >= 2 {
		_, err = time.Parse("2006-01-02", args[0])
		if err != nil {
			return nil, fmt.Errorf("Failed to parse birthday: %s", err)
		}
		birthday = args[0]

		age, err = strconv.Atoi(args[1])
		if err != nil {
			return nil, fmt.Errorf("Failed to convert year to int: %s", err)
		}
	}

	if len(args) > 2 {
		// parse sex
	}

	if len(args) > 3 {
		// parse features
	}

	return &v1.PersonRequest{
		Api:  "v1",
		Lang: "ru-RU",

		PersonProfile: &v1.PersonProfile{
			Name:     "A Person",
			Gender:   1,
			Features: 0,
			Birthday: birthday,
			Age:      int32(age),
		},
	}, nil
}

// func CommandNewHelp() {}
