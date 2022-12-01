package main

import (
	"log"
	"voidmanager/bot"
	"voidmanager/db"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config, err := utils.LoadConfig(".", "app")
	if err != nil {
		log.Fatalf("[ERROR] unable to load config file, %v", err.Error())
	}

	dbClient := db.InitDB(&config)

	benny := bot.New(&config, dbClient)
	benny.SetupIntents(discordgo.IntentsGuildMessages)
	benny.AddHandlers()
	benny.StartBot()
}
