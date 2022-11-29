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

	command, args := splitMessageContent(m.Content)

	if command == "!ping" {
		commands.PingCommand(s, m, args)
	}
}

func splitMessageContent(content string) (prefix string, args []string) {
	splitContent := strings.Split(content, " ")

	prefix = splitContent[0]
	args = splitContent[1:]

	return prefix, args
}
