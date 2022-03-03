package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	//"os/signal"
	"strings"
	//"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Token := os.Getenv("TOKEN")

	dg, err := discordgo.New("Bot " + Token)

	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsAll

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Void Manager bot is ready.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	var command = strings.Split(m.Content, " ")

	if command[0] == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")

		if err != nil {
			fmt.Println(err)
		}
	}
}
