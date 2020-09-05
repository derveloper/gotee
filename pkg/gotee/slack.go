package gotee

import (
	"github.com/slack-go/slack"
	"log"
	"os"
	"strings"
)

func SlackOutput(token string, channel string) func(lines []string) {
	api := slack.New(
		token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
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
			} else {
				log.Println("updated message")
			}
		} else {
			_, ts, _, err := api.SendMessage(channel,
				slack.MsgOptionText(msg, false),
				slack.MsgOptionAsUser(true))
			if err != nil {
				log.Println(err)
			}
			timestamp = ts
			log.Printf("posted new message %s\n", timestamp)
		}
	}
	return lineFunc
}
