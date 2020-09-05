package main

import (
	"gotee/pkg/gotee"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigGoTee struct {
	Token   string `yaml:"slack_token" env:"GOTEE_SLACK_TOKEN"`
	Channel string `yaml:"slack_channel" env:"GOTEE_SLACK_CHANNEL"`
}

func main() {
	var cfg ConfigGoTee

	configFile := "/etc/gotee/config.yml"

	info, err := os.Stat(configFile)
	if !os.IsNotExist(err) && !info.IsDir() {
		err = cleanenv.ReadConfig(configFile, &cfg)
		if err != nil {
			log.Println(err)
		}
	} else {
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Println(err)
		}
	}

	lineFunc := gotee.SlackOutput(cfg.Token, cfg.Channel)

	gotee.Tee(lineFunc, 5*time.Second)
}