package events

import (
	"strings"
	"voidmanager/bot/commands"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, _ := splitMessageContent(m.Content)

	if command == "!ping" {
		commands.PingCommand(s, m)
	}
}

func splitMessageContent(content string) (prefix string, values []string) {
	splitContent := strings.Split(content, " ")

	prefix = splitContent[0]
	values = splitContent[1:]

	return
}
