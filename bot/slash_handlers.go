package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var commandHandlers InteractionCommandHandlersMap = InteractionCommandHandlersMap{
	"shoutout": SlashPingCommand,
	"hello":    SlashHelloCommand,
}

func (h *Handler) SlashCommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println(i.ApplicationCommandData().Name, i.Member.User.ID, i.Member.User)
	if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}

func SlashPingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	}
}

func SlashHelloCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
	}
}
