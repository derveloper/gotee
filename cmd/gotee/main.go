package main

import (
	"flag"
	"fmt"
	"gotee/pkg/gotee"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/slack-go/slack"
)

type ConfigGoTee struct {
	Token   string `yaml:"slack_token" env:"GOTEE_SLACK_TOKEN" env-default:""`
	Channel string `yaml:"slack_channel" env:"GOTEE_SLACK_CHANNEL" env-default:""`
}

func main() {
	var cfg ConfigGoTee

	err := cleanenv.ReadConfig("/etc/gotee/config.yml", &cfg)
	if err != nil {
		log.Println(err)
	}

	output := flag.String("output", "slack", "output method to use")
	flag.Parse()

	lineFunc := func(lines []string) {
		fmt.Print(lines)
	}

	switch *output {
	case "slack":
		api := slack.New(
			cfg.Token,
			slack.OptionDebug(true),
			slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
		)
		var timestamp string
		lines := make([]string, 0)
		lineFunc = func(newLines []string) {
			lines = append(lines, newLines...)
			msg := strings.Join(lines, "\n")
			if timestamp != "" {
				_, _, _, err := api.UpdateMessage(cfg.Channel, timestamp, slack.MsgOptionText(msg, false))
				if err != nil {
					log.Println(err)
				} else {
					log.Println("updated message")
				}
			} else {
				_, ts, _, err := api.SendMessage(cfg.Channel,
					slack.MsgOptionText(msg, false),
					slack.MsgOptionAsUser(true))
				if err != nil {
					log.Println(err)
				}
				timestamp = ts
				log.Printf("posted new message %s\n", timestamp)
			}
		}
		break
	}

	gotee.Tee(lineFunc, 5*time.Second)
}
