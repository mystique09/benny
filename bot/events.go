package bot

import (
	"context"
	"fmt"
	"log"
	"strings"
	"voidmanager/db/ent"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) Ready(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("\nLogged in as %v \n", r.User.Username)
	fmt.Println("Token: <redacted>")
	fmt.Printf("Version: %v \n", h.cfg.BotVersion)
	fmt.Printf("Onwer: %v \n", h.cfg.BotOwner)
	fmt.Printf("Application ID: %v \n", h.cfg.BotApplicationId)
	fmt.Printf("Guild ID: %v \n\n", h.cfg.BotGuildId)
}

func (h *Handler) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, args := splitMessageContent(m.Content)

	if strings.HasPrefix(command, "!") {
		log.Printf("user <@%v> used the %v command with args %v", m.Author.ID, command, args)
	}

	if command == "!ping" {
		h.PingCommand(s, m, args)
	}
}

func (h *Handler) MemberCreate(s *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	_, err := h.client.User.
		Create().
		SetID(newMember.User.ID).
		SetName(newMember.User.Username).
		Save(context.Background())

	if err != nil {
		log.Printf("user %v already exist", newMember.User.ID)
	}
}

func (h *Handler) MemberRemove(s *discordgo.Session, member *discordgo.GuildMemberRemove) {
	if err := h.client.User.DeleteOne(&ent.User{
		ID:   member.User.ID,
		Name: member.User.Username,
	}).Exec(context.Background()); err != nil {
		log.Printf("user %v doesn't exist in database", member.User.ID)
	}
}

func splitMessageContent(content string) (prefix string, args []string) {
	splitContent := strings.Split(content, " ")

	prefix = splitContent[0]
	args = splitContent[1:]

	return prefix, args
}
