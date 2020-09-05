package gotee

import (
	"github.com/slack-go/slack"
	"log"
	"strings"
)

func SlackOutput(token string, channel string) func(lines []string) {
	api := slack.New(
		token,
	)
	var timestamp string
	lines := make([]string, 0)
	lineFunc := func(newLines []string) {
		lines = append(lines, newLines...)
		msg := strings.Join(lines, "\n")
		if timestamp != "" {
			_, _, _, err := api.UpdateMessage(channel, timestamp, slack.MsgOptionText(msg, false))
			if err != nil {
				log.Println(err)
			}
		} else {
			_, ts, _, err := api.SendMessage(channel,
				slack.MsgOptionText(msg, false),
				slack.MsgOptionAsUser(true))
			if err != nil {
				log.Println(err)
			}
			timestamp = ts
		}
	}
	return lineFunc
}
