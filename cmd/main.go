package main

import (
	"flag"
	"log"
	"voidmanager/bot"
	"voidmanager/db"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	removeCommandsAfterShutdown = flag.Bool("rmslash", true, "Remove all registered slash commands after shutdown.")
)

func main() {
	config, err := utils.LoadConfig(".", "app")
	if err != nil {
		log.Fatalf("[ERROR] unable to load config file, %v", err.Error())
	}

	dbClient := db.InitDB(&config)

	benny := bot.NewBot(&config, dbClient)
	benny.SetupIntents(discordgo.IntentsGuildMessages)
	benny.AddHandlers()
	benny.StartBot(*removeCommandsAfterShutdown)
}
