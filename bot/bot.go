package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"voidmanager/db/ent"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	dg      *discordgo.Session
	handler Handler
}

func New(cfg *utils.Config, client *ent.Client) *Bot {
	dg, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	handler := Handler{
		cfg,
		client,
	}

	return &Bot{
		dg:      dg,
		handler: handler,
	}
}

func (bot *Bot) SetupIntents(intents discordgo.Intent) {
	bot.dg.Identify.Intents = intents
}

func (bot *Bot) AddHandlers() {
	bot.dg.AddHandler(bot.handler.Ready)
	bot.dg.AddHandler(bot.handler.MessageCreate)
	bot.dg.AddHandler(bot.handler.MemberCreate)
	bot.dg.AddHandler(bot.handler.MemberRemove)
}

func (bot *Bot) StartBot() {
	if err := bot.dg.Open(); err != nil {
		log.Fatal(err.Error())
	}
	bot.listenShutdown()
}

func (bot *Bot) listenShutdown() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := bot.dg.Close(); err != nil {
		log.Fatal(err.Error())
	}
}
