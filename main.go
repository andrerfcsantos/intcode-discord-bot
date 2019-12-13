package main

import (
	"flag"
	"fmt"
	"intcode-discord-bot/handlers"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	dgo "github.com/bwmarrin/discordgo"
)

// BOT_TOKEN represents the discord authentication token
var BOT_TOKEN string

var session *dgo.Session

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {

	// Discord Authentication Token
	BOT_TOKEN = os.Getenv("INTCODE_DISCORD_BOT_TOKEN")
	if BOT_TOKEN == "" {
		flag.StringVar(&BOT_TOKEN, "bot-token", "", "Discord Authentication Token")
	}
	flag.Parse()
}

func main() {
	var err error

	err = setOutput()

	if err != nil {
		fmt.Printf("error setting up logs: %v", err)
		return
	}

	if BOT_TOKEN == "" {
		log.Print("No bot token specified. Please specify one using the INTCODE_DISCORD_BOT_TOKEN environment variable or the -bot-token flag.")
		return
	}

	session, err = dgo.New("Bot " + BOT_TOKEN)
	if err != nil {
		log.Printf("error getting new session: %v", err)
		return
	}

	err = session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		return
	}
	defer session.Close()

	session.UpdateStatus(0, "!intcode help")
	session.AddHandler(handlers.CreateMessage)

	// Wait for a CTRL-C
	log.Printf("Intcode VM running! (CTRL-C to stop)")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}

func setOutput() error {
	err := os.MkdirAll("/var/log/intcode/", os.FileMode(0755))
	if err != nil {
		return err
	}

	file, err := os.OpenFile("/var/log/intcode/bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	return nil
}
