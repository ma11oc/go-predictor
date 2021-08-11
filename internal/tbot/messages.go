package tbot

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	pb "github.com/ma11oc/go-predictor/pkg/api/v1"
)

type header struct {
	Title   string
	Content string
}

type section struct {
	Title   string
	Content string
}

type message struct {
	Header   *header
	Sections []*section
	Footer   string
}

var msgTemplate = `
{{ with .Header }}
{{- with .Title }}
<b>{{ . }}</b>
{{ end }}
{{- with .Content }}
{{ . }}
{{- end }}
{{- end }}

{{- range .Sections }}
{{- with .Title }}
<b>{{ . }}</b>
{{- end }}
{{- with .Content }}
<pre>
{{ . }}
</pre>
{{- end }}
{{- end }}

{{- with .Footer }}
<u>{{ . }}</u>
{{- end }}
`

// MakeMessageByCallback handles callbacks
func MakeMessageByCallback(p *pb.Person, callback string) (string, error) {
	m := &message{
		Header: &header{
			Title: fmt.Sprintf("%s, %d, %s%s, %s",
				p.GetName(),
				p.GetAge(),
				p.GetBaseCards()["main"].GetRank(),
				p.GetBaseCards()["main"].GetSuit(),
				p.GetBirthday(),
			),
		},
	}

	cb := strings.Split(callback, ":")
	_, cbType, cbSet, cbCard, cbQuery := cb[0], cb[1], cb[2], cb[3], cb[4]

	switch cbType {
	// card:base:result:meta
	// card:base:result:desc
	case "card":
		switch cbSet {
		case "base":
			switch cbQuery {
			case "desc":
				m.Sections = append(m.Sections, &section{
					Title:   "Card: " + p.GetBaseCards()[cbCard].GetRank() + p.GetBaseCards()[cbCard].GetSuit(),
					Content: "",
				})

				m.Sections = append(m.Sections, &section{
					Title:   "Keywords:",
					Content: p.GetBaseCards()[cbCard].Meaning.Keywords,
				})

				m.Sections = append(m.Sections, &section{
					Title:   "Description:",
					Content: p.GetBaseCards()[cbCard].Meaning.Description,
				})
			case "meta":
				m.Sections = append(m.Sections, &section{
					Title:   cbCard + ":",
					Content: p.GetBaseCards()[cbCard].XMeta,
				})
			default:
				return "", nil
			}
		case "pers":
			return "", nil
		case "karm":
			return "", nil
		default:
			return "", nil
		}
	// planet:all:period:meta
	// planet:%s:horizontal:desc
	case "planet":
		switch cbSet {
		case "all":
			m.Sections = append(m.Sections, &section{
				Title:   cbCard + ":",
				Content: p.GetPlanetCycles()["mars"].GetCards()[cbCard].XMeta,
			})
		case "mercury", "venus", "mars", "jupiter", "saturn", "uranus", "neptune":
			m.Sections = append(m.Sections, &section{
				Title:   "Card: " + p.GetPlanetCycles()[cbSet].GetCards()[cbCard].GetRank() + p.GetPlanetCycles()[cbSet].GetCards()[cbCard].GetSuit(),
				Content: "",
			})

			m.Sections = append(m.Sections, &section{
				Title:   "Keywords:",
				Content: p.GetPlanetCycles()[cbSet].GetCards()[cbCard].Meaning.Keywords,
			})

			m.Sections = append(m.Sections, &section{
				Title:   "Description:",
				Content: p.GetPlanetCycles()[cbSet].GetCards()[cbCard].Meaning.Description,
			})
		default:
			return "", nil
		}
	default:
		return "", nil
	}

	return renderMessage(msgTemplate, m)
}

// MakeErrorMessage prepares error message to be send to the user
func MakeErrorMessage(title string, e error) (*tgbotapi.MessageConfig, error) {
	m := &message{
		Sections: []*section{
			{
				Title:   title,
				Content: fmt.Sprint(e),
			},
		},
	}

	text, err := renderMessage(msgTemplate, m)
	if err != nil {
		return nil, err
	}

	msg := &tgbotapi.MessageConfig{
		Text:      text,
		ParseMode: tgbotapi.ModeHTML,
	}

	return msg, nil
}

func renderMessage(tmpl string, data interface{}) (string, error) {
	t, err := template.New("msg").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil

}
