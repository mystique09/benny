package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type InteractionCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate)
type InteractionCommandHandlersMap = map[string]InteractionCommandHandler

var commands []*discordgo.ApplicationCommand = []*discordgo.ApplicationCommand{
	{
		Name:        "shoutout",
		Description: "Shoutout the mentioned user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "the user to shoutout",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    false,
			},
		},
	},
	{
		Name:        "hello",
		Description: "Says hello to you!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "the user to shoutout",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    false,
			},
		},
	},
}

func (bot *Bot) createSlashCommands(registeredCommands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	log.Println("Creating slash commands...")

	for i, cmd := range commands {
		log.Printf("Adding `%v` slash command", cmd.Name)
		command, err := bot.dg.ApplicationCommandCreate(bot.dg.State.User.ID, bot.handler.cfg.BotGuildId, cmd)
		if err != nil {
			log.Fatalf("[ERROR] cannot create %v slash command: %v ", cmd.Name, err.Error())
		}
		registeredCommands[i] = command
	}

	log.Printf("Total slash commands created: %d", len(registeredCommands))
	return registeredCommands
}

func (bot *Bot) removeSlashCommands(registeredCommands []*discordgo.ApplicationCommand) {
	log.Println("Removing slash commands...")

	for _, cmd := range registeredCommands {
		log.Printf("Removing `%v` slash command", cmd.Name)
		err := bot.dg.ApplicationCommandDelete(bot.dg.State.User.ID, bot.handler.cfg.BotGuildId, cmd.ID)
		if err != nil {
			log.Fatalf("[ERROR] cannot delete %v slash command %v", cmd.Name, err.Error())
		}
	}

	log.Printf("Total slash commands removed: %d", len(registeredCommands))
}

var commandHandlers InteractionCommandHandlersMap = InteractionCommandHandlersMap{
	"shoutout": SlashPingCommand,
	"hello":    SlashHelloCommand,
}

func (h *Handler) SlashCommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println(i.ApplicationCommandData().Name, i.Member.User.ID, i.Member.User)
	if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}