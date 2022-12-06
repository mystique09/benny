package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type InteractionCommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate, h *Handler)
type InteractionCommandHandlersMap = map[string]InteractionCommandHandler

var slashCommandHandlers InteractionCommandHandlersMap = InteractionCommandHandlersMap{
	"shoutout": ShoutoutHandler,
	"hello":    HelloHandler,
	"ping":     PingHandler,
}

func (bot *Bot) SlashCommandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println(i.ApplicationCommandData().Name, i.Member.User.ID, i.Member.User)
	if handler, ok := slashCommandHandlers[i.ApplicationCommandData().Name]; ok {
		handler(s, i, &bot.handler)
	}
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate, h *Handler) {
	botLatency := s.HeartbeatLatency()

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Bot latency is %d ms", botLatency.Milliseconds()),
		},
	}); err != nil {
		_, err := s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "Something went wrong",
		})

		log.Println(err.Error())
		return
	}
}

func HelloHandler(s *discordgo.Session, i *discordgo.InteractionCreate, h *Handler) {
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

func ShoutoutHandler(s *discordgo.Session, i *discordgo.InteractionCreate, h *Handler) {
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
