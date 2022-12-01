package bot

import (
	"log"
	"voidmanager/db/ent"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	cfg    *utils.Config
	client *ent.Client
}

func (h *Handler) PingCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if _, err := s.ChannelMessageSend(m.ChannelID, "Pong!"); err != nil {
		log.Println(err.Error())
	}
}
