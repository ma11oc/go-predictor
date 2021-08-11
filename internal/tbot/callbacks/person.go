package tbot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"

	tbot "github.com/ma11oc/go-predictor/internal/tbot"
	model "github.com/ma11oc/go-predictor/internal/tbot/db"
	markups "github.com/ma11oc/go-predictor/internal/tbot/markups"
	v1 "github.com/ma11oc/go-predictor/pkg/api/v1"
)

// HandlePersonCallbackQuery answers to person's callback
func HandlePersonCallbackQuery(ctx context.Context, query *tgbotapi.CallbackQuery, psc v1.PredictorServiceClient, db *gorm.DB, bot *tgbotapi.BotAPI) error {
	var cspan opentracing.Span

	pspan, _ := opentracing.StartSpanFromContext(ctx, "HandlePersonCallbackQuery")
	defer pspan.Finish()

	if _, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(query.ID, query.Data)); err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
	}

	person := &model.PersonProfile{}
	result := db.Debug().Model(&model.PersonProfile{}).
		Select("person_profiles.name, person_profiles.gender, person_profiles.birthday, person_profiles.age, person_profiles.features, requests.chat_id AS request_chat_id, requests.message_id, responses.message_id AS response_message_id").
		Joins("LEFT JOIN requests ON person_profiles.request_id = requests.id").
		Joins("LEFT JOIN responses ON responses.request_id = requests.id").
		Where("request_chat_id = ? AND response_message_id = ?", query.Message.Chat.ID, query.Message.MessageID).
		Scan(person)

	if result.RowsAffected == 0 {
		err := fmt.Errorf("Failed to get PersonProfile")
		pspan.SetTag("error", true).LogFields(log.Error(err))
		// TODO: send error message
		return err
	}

	if result.Error != nil {
		pspan.SetTag("error", true).LogFields(log.Error(result.Error))
		return result.Error
	}

	// TODO: add api version and lang to Request

	pspan.LogFields(log.String("person", fmt.Sprintf("%+v", person.PersonProfile)))

	preq := &v1.PersonRequest{
		Api:  "v1",
		Lang: "ru-RU",

		PersonProfile: &v1.PersonProfile{
			Name:     person.PersonProfile.Name,
			Age:      person.PersonProfile.Age,
			Gender:   person.PersonProfile.Gender,
			Birthday: person.PersonProfile.Birthday,
			Features: person.PersonProfile.Features,
		},
	}

	presp, err := psc.ComputePerson(opentracing.ContextWithSpan(ctx, pspan), preq)
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	text, err := tbot.MakeMessageByCallback(presp.Person, query.Data)
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	msg := tgbotapi.NewEditMessageText(
		query.Message.Chat.ID,
		query.Message.MessageID,
		text,
	)
	msg.ParseMode = tgbotapi.ModeHTML

	markup, err := markups.MakePersonMarkup(presp.Person)
	if err != nil {
		pspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	msg.ReplyMarkup = &markup

	cspan = opentracing.StartSpan("bot.Send", opentracing.ChildOf(pspan.Context()))
	defer cspan.Finish()

	if _, err := bot.Send(msg); err != nil {
		cspan.SetTag("error", true).LogFields(log.Error(err))
		return err
	}

	return nil
}
