package bot

import (
	"fmt"
	"log"
	"voidmanager/db/ent"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

type Prefix = string
type GuildId = string

type Commands struct {
	client   *ent.Client
	cfg      *utils.Config
	prefixes map[GuildId]Prefix
}

type CommandHandler = func(s *discordgo.Session, m *discordgo.MessageCreate, args []string, c *Commands)

var commandHandlers map[string]CommandHandler = map[string]CommandHandler{
	"ping": pingCommand,
}

func pingCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string, c *Commands) {
	pingMsg := fmt.Sprintf("Heartbeat latency is %d ms", s.HeartbeatLatency().Milliseconds())
	if _, err := s.ChannelMessageSend(m.ChannelID, pingMsg); err != nil {
		log.Println(err.Error())
	}
}
