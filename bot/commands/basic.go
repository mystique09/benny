package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "Pong!"); err != nil {
		log.Println(err.Error())
	}
}
