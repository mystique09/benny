package bot

import (
	"context"
	"fmt"
	"log"
	"voidmanager/db/ent"

	"github.com/bwmarrin/discordgo"
)

var slashCommandHandlers InteractionCommandHandlersMap = InteractionCommandHandlersMap{
	"shoutout":   SlashPingCommand,
	"hello":      SlashHelloCommand,
	"set-prefix": SlashSetPrefixCommand,
	"prefix":     SlashGetBotPrefixCommand,
}

func (e *Events) SlashCommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println(i.ApplicationCommandData().Name, i.Member.User.ID, i.Member.User)
	if handler, ok := slashCommandHandlers[i.ApplicationCommandData().Name]; ok {
		handler(s, i, e.commands)
	}
}

func SlashPingCommand(s *discordgo.Session, i *discordgo.InteractionCreate, c *Commands) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("shoutout to <@%v>", optionMap["user"].Value),
		},
	}); err != nil {
		_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong",
		})

		log.Println(err.Error())
		return
	}
}

func SlashHelloCommand(s *discordgo.Session, i *discordgo.InteractionCreate, c *Commands) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("hello <@%v>", optionMap["user"].Value),
		},
	}); err != nil {
		_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong",
		})

		log.Println(err.Error())
		return
	}
}

func SlashSetPrefixCommand(s *discordgo.Session, i *discordgo.InteractionCreate, c *Commands) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	newPrefix := optionMap["prefix"].StringValue()

	checkGuild, err := c.client.Guild.Get(context.Background(), i.GuildID)
	if err != nil || checkGuild == nil {
		_, err := c.client.Guild.
			Create().
			SetID(i.GuildID).
			SetBotPrefix(newPrefix).
			Save(context.Background())
		if err != nil {
			log.Printf("unable to create guild: %s", err.Error())
			if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Unable to set prefix, something went wrong.",
				},
			}); err != nil {
				_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "Something went wrong",
				})

				log.Println(err.Error())
				return
			}
			return
		}
	}

	updatedGuild, err := c.client.Guild.UpdateOne(&ent.Guild{
		ID:        i.GuildID,
		BotPrefix: checkGuild.BotPrefix,
	}).SetBotPrefix(newPrefix).
		Save(context.Background())
	if err != nil {
		log.Println(err.Error())

		if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Unable to set prefix, something went wrong.",
			},
		}); err != nil {
			_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "Something went wrong",
			})

			log.Println(err.Error())
			return
		}

		return
	}

	// need to set the prefix in the memory
	c.prefixes[i.GuildID] = updatedGuild.BotPrefix
	log.Printf(`The %s guild/server, sets a new prefix, new bot prefix is %s.`, i.GuildID, updatedGuild.BotPrefix)

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(`the new bot prefix is %s.`, updatedGuild.BotPrefix),
		},
	}); err != nil {
		_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong",
		})

		log.Println(err.Error())
		return
	}
}

func SlashGetBotPrefixCommand(s *discordgo.Session, i *discordgo.InteractionCreate, c *Commands) {
	currentBotPrefix, ok := c.prefixes[i.GuildID]
	if !ok {
		log.Printf("guild %s doesn't have a prefix in cache", i.GuildID)
	}

	msg := fmt.Sprintf("the current guild prefix is %s", currentBotPrefix)
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	}); err != nil {
		_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong",
		})

		log.Println(err.Error())
		return
	}
}
