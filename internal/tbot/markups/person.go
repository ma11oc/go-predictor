package tbot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	core "github.com/ma11oc/go-predictor/internal/core"
	pb "github.com/ma11oc/go-predictor/pkg/api/v1"
)

func MakePersonMarkup(p *pb.Person) (tgbotapi.InlineKeyboardMarkup, error) {

	mc := p.GetBaseCards()["main"]
	sc := p.GetBaseCards()["source"]
	dc := p.GetBaseCards()["drain"]
	lc := p.GetBaseCards()["longterm"]
	pc := p.GetBaseCards()["pluto"]
	rc := p.GetBaseCards()["result"]
	// kcc := p.GetKarmaCards()

	markup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Base cards", "title_base_cards")),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("source", "person:card:base:source:meta"),
			tgbotapi.NewInlineKeyboardButtonData("main", "person:card:base:main:meta"),
			tgbotapi.NewInlineKeyboardButtonData("drain", "person:card:base:drain:meta"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(sc.Rank+sc.Suit, "person:card:base:source:desc"),
			tgbotapi.NewInlineKeyboardButtonData(mc.Rank+mc.Suit, "person:card:base:main:desc"),
			tgbotapi.NewInlineKeyboardButtonData(dc.Rank+dc.Suit, "person:card:base:drain:desc"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("longterm", "person:card:base:longterm:meta"),
			tgbotapi.NewInlineKeyboardButtonData("pluto", "person:card:base:pluto:meta"),
			tgbotapi.NewInlineKeyboardButtonData("result", "person:card:base:result:meta"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(lc.Rank+lc.Suit, "person:card:base:longterm:desc"),
			tgbotapi.NewInlineKeyboardButtonData(pc.Rank+pc.Suit, "person:card:base:pluto:desc"),
			tgbotapi.NewInlineKeyboardButtonData(rc.Rank+rc.Suit, "person:card:base:result:desc"),
		),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Planet Cycles", "planet_cycles_title")),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("period", "person:planet:noop:period:meta"),
			tgbotapi.NewInlineKeyboardButtonData("horizontal", "person:planet:all:horizontal:meta"),
			tgbotapi.NewInlineKeyboardButtonData("vertical", "person:planet:all:vertical:meta"),
		),
	)

	pcc := p.GetPlanetCycles()

	for _, planet := range core.PlanetsOrder {
		markup.InlineKeyboard = append(
			markup.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf(
						"%02d.%02d-%02d.%02d | %s",
						pcc[planet].GetStart().GetDay(),
						pcc[planet].GetStart().GetMonth(),
						pcc[planet].GetEnd().GetDay(),
						pcc[planet].GetEnd().GetMonth(),
						pcc[planet].GetPlanet().GetSymbol(),
					),
					fmt.Sprintf("person:planet:%s:cycle:meta", planet),
				),
				tgbotapi.NewInlineKeyboardButtonData(
					pcc[planet].GetCards()["horizontal"].GetRank()+pcc[planet].GetCards()["horizontal"].GetSuit(),
					fmt.Sprintf("person:planet:%s:horizontal:desc", planet),
				),
				tgbotapi.NewInlineKeyboardButtonData(
					pcc[planet].GetCards()["vertical"].GetRank()+pcc[planet].GetCards()["vertical"].GetSuit(),
					fmt.Sprintf("person:planet:%s:vertical:desc", planet),
				),
			),
		)
	}

	return markup, nil
}
