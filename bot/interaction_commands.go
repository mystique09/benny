package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var slashCommands []*discordgo.ApplicationCommand = []*discordgo.ApplicationCommand{
	{
		Name:        "shoutout",
		Description: "Shoutout the mentioned user",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "the user to shoutout",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
		},
	},
	{
		Name:        "hello",
		Description: "Says hello to you!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "user",
				Description: "the user to say hello",
				Type:        discordgo.ApplicationCommandOptionUser,
				Required:    true,
			},
		},
	},
	{
		Name:        "ping",
		Description: "get the current bot latency",
	},
}

func (bot *Bot) createSlashCommands(registeredCommands []*discordgo.ApplicationCommand) []*discordgo.ApplicationCommand {
	log.Println("Creating slash commands...")

	for i, cmd := range slashCommands {
		log.Printf("Adding `%v` slash command", cmd.Name)
		command, err := bot.dg.ApplicationCommandCreate(bot.dg.State.User.ID, bot.handler.cfg.BotGuildId, cmd)
		if err != nil {
			log.Fatalf("[ERROR] cannot create %v slash command: %v ", cmd.Name, err.Error())
		}
		registeredCommands[i] = command
	}

	log.Printf("Successfully added all slash commands to guild: %s", bot.handler.cfg.BotGuildId)
	log.Printf("[SUCCESS] Total slash commands created: %d", len(registeredCommands))
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
