package bot

import (
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h *Handler) Ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("\nLogged in as %v \n", r.User.Username)
	log.Printf("Version: %v \n", h.cfg.BotVersion)
	log.Printf("Onwer: %v \n", h.cfg.BotOwner)
	log.Printf("Application ID: %v \n", h.cfg.BotApplicationId)
	log.Printf("Guild ID: %v \n", h.cfg.BotGuildId)
	log.Printf("Total slash commands: %d\n\n", len(slashCommands))
}

func (h *Handler) GuildCreate(s *discordgo.Session, newGuild *discordgo.GuildCreate) {
	_, err := h.client.Guild.
		Create().
		SetID(newGuild.ID).
		SetBotPrefix("!").
		Save(context.Background())

	if err != nil {
		log.Println("guild already exist")
	}

	log.Printf(`New guild added %s`, newGuild.ID)
}

func (h *Handler) GuildRemove(s *discordgo.Session, guild *discordgo.GuildDelete) {
	if err := h.client.Guild.DeleteOneID(guild.ID).Exec(context.Background()); err != nil {
		log.Printf("guild %v doesn't exist in database", guild.ID)
	}
}

func (h *Handler) MemberCreate(s *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	_, err := h.client.User.
		Create().
		SetID(newMember.Member.User.ID).
		SetName(newMember.Member.User.Username).
		Save(context.Background())

	if err != nil {
		log.Printf("user %v already exist", newMember.User.ID)
	}

	log.Printf(`New member %s added`, newMember.Member.User.ID)
}

func (h *Handler) MemberRemove(s *discordgo.Session, member *discordgo.GuildMemberRemove) {
	if err := h.client.User.DeleteOneID(member.Member.User.ID).Exec(context.Background()); err != nil {
		log.Printf("user %v doesn't exist in database", member.Member.User.ID)
	}
}
