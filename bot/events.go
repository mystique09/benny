package bot

import (
	"context"
	"fmt"
	"log"
	"strings"
	"voidmanager/db/ent"
	"voidmanager/utils"

	"github.com/bwmarrin/discordgo"
)

type Events struct {
	client   *ent.Client
	cfg      *utils.Config
	commands *Commands
}

func (e *Events) Ready(s *discordgo.Session, r *discordgo.Ready) {
	allGuilds, err := e.client.Guild.Query().All(context.Background())

	for _, guild := range allGuilds {
		e.commands.prefixes[guild.ID] = guild.BotPrefix
	}

	log.Printf("Added prefixes %v", e.commands.prefixes)

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("\nLogged in as %v \n", r.User.Username)
	fmt.Println("Token: <redacted>")
	fmt.Printf("Version: %v \n", e.cfg.BotVersion)
	fmt.Printf("Onwer: %v \n", e.cfg.BotOwner)
	fmt.Printf("Application ID: %v \n", e.cfg.BotApplicationId)
	fmt.Printf("Guild ID: %v \n", e.cfg.BotGuildId)
	fmt.Printf("Total commands: %d\n", len(commandHandlers))
	fmt.Printf("Total slash commands: %d\n\n", len(slashCommandHandlers))
}

func (e *Events) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, args := splitMessageContent(m.Content)

	guildPrefix, ok := e.commands.prefixes[m.GuildID]
	if !ok {
		if strings.HasPrefix(command, "!") {
			command = strings.TrimPrefix(command, "!")
			log.Printf("user <@%v> used the %v command with args %v", m.Author.ID, command, args)

			cmd, ok := commandHandlers[command]
			if !ok {
				log.Println("command doesn't exist")
				return
			}
			cmd(s, m, args, e.commands)
		}

		return
	}

	if strings.HasPrefix(command, guildPrefix) {
		command = strings.TrimPrefix(command, guildPrefix)
		log.Printf("user <@%v> used the %v command with args %v", m.Author.ID, command, args)

		cmd, ok := commandHandlers[command]
		if !ok {
			log.Println("command doesn't exist")
			return
		}
		cmd(s, m, args, e.commands)
	}
}

func (e *Events) GuildCreate(s *discordgo.Session, newGuild *discordgo.GuildCreate) {
	_, err := e.client.Guild.
		Create().
		SetID(newGuild.ID).
		SetBotPrefix("!").
		Save(context.Background())

	if err != nil {
		log.Println("guild already exist")
	}

	log.Printf(`New guild %s added`, newGuild.ID)
}

func (e *Events) GuildRemove(s *discordgo.Session, guild *discordgo.GuildDelete) {
	if err := e.client.Guild.DeleteOneID(guild.ID).Exec(context.Background()); err != nil {
		log.Printf("guild %v doesn't exist in database", guild.ID)
	}
}

func (e *Events) MemberCreate(s *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	_, err := e.client.User.
		Create().
		SetID(newMember.Member.User.ID).
		SetName(newMember.Member.User.Username).
		Save(context.Background())

	if err != nil {
		log.Printf("user %v already exist", newMember.User.ID)
	}

	log.Printf(`New member %s added`, newMember.Member.User.ID)
}

func (e *Events) MemberRemove(s *discordgo.Session, member *discordgo.GuildMemberRemove) {
	if err := e.client.User.DeleteOneID(member.Member.User.ID).Exec(context.Background()); err != nil {
		log.Printf("user %v doesn't exist in database", member.Member.User.ID)
	}
}

func splitMessageContent(content string) (command string, args []string) {
	splitContent := strings.Split(content, " ")

	command = splitContent[0]
	args = splitContent[1:]

	return command, args
}
