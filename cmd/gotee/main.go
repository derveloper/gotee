package main

import (
	"gotee/pkg/gotee"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	flag "github.com/spf13/pflag"
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

	flag.StringVarP(&cfg.Channel, "channel", "c", cfg.Channel, "set the channel to post to")
	showHelp := flag.BoolP("help", "h", false, "print this help")
	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	lineFunc := gotee.SlackOutput(cfg.Token, cfg.Channel)

	gotee.Tee(lineFunc, 5*time.Second)
}
