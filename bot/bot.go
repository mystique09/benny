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

type Handler struct {
	cfg    *utils.Config
	client *ent.Client
}

func NewBot(cfg *utils.Config, client *ent.Client) *Bot {
	dg, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	return &Bot{
		dg: dg,
		handler: Handler{
			cfg:    cfg,
			client: client,
		},
	}
}

func (bot *Bot) SetupIntents(intents discordgo.Intent) {
	bot.dg.Identify.Intents = intents
}

func (bot *Bot) AddHandlers() {
	bot.dg.AddHandler(bot.handler.Ready)
	bot.dg.AddHandler(bot.handler.MemberCreate)
	bot.dg.AddHandler(bot.handler.MemberRemove)
	bot.dg.AddHandler(bot.SlashCommandsHandler)
}

func (bot *Bot) StartBot(removeCommandsAfterShutdown bool) {
	if err := bot.dg.Open(); err != nil {
		log.Fatal(err.Error())
	}

	var registeredCommands []*discordgo.ApplicationCommand = make([]*discordgo.ApplicationCommand, len(slashCommands))
	bot.createSlashCommands(registeredCommands)

	defer bot.dg.Close()

	bot.listenShutdown()

	if removeCommandsAfterShutdown {
		bot.removeSlashCommands(registeredCommands)
	}

	log.Println("Gracefully shutting down...")
	log.Println("Shutdown successful.")
}

func (bot *Bot) listenShutdown() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if err := bot.dg.Close(); err != nil {
		log.Fatal(err.Error())
	}
}
