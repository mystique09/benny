package main

import (
	"log"
	"voidmanager/bot"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config, err := utils.LoadConfig(".", "app")
	if err != nil {
		log.Fatalf("[ERROR] unable to load config file, %v", err.Error())
	}

	benny := bot.New(&config)
	benny.SetupIntents(discordgo.IntentsGuildMessages)
	benny.AddHandlers()
	benny.StartBot()
}
