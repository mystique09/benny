package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"voidmanager/bot/events"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	cfg *utils.Config
	dg  *discordgo.Session
}

func New(cfg *utils.Config) *Bot {
	dg, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}
	return &Bot{cfg, dg}
}

func (bot *Bot) SetupIntents(intents discordgo.Intent) {
	bot.dg.Identify.Intents = intents
}

func (bot *Bot) AddHandlers() {
	bot.dg.AddHandler(events.MessageCreate)
}

func (bot *Bot) StartBot() {
	if err := bot.dg.Open(); err != nil {
		log.Fatal(err.Error())
	}
	bot.printBanner()
	bot.listenShutdown()
}

func (bot *Bot) printBanner() {
	fmt.Printf(`
==== Bot started ====
Bot Version: %v
Bot Token: %v
Bot Onwer: %v
Channel ID: %v
Application ID: %v
`,
		bot.cfg.BotVersion,
		bot.cfg.BotToken,
		bot.cfg.BotOwner,
		bot.cfg.BotChannelId,
		bot.cfg.BotApplicationId)
}

func (bot *Bot) listenShutdown() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := bot.dg.Close(); err != nil {
		log.Fatal(err.Error())
	}
}
