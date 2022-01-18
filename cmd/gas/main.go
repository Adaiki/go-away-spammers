package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const Version = "v0.0.0-alpha"

var Session *discordgo.Session

func init() {

	godotenv.Load(".env")
	Session, _ = discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	Session.SyncEvents = false
}

func main() {

	if Session.Token == "" {
		log.Println("You must provide a Discord authentication token.")
		return
	}

	Session.AddHandler(messageCreate)
	Session.Identify.Intents = discordgo.IntentsGuildMessages

	err := Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	Session.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// TODO
}
