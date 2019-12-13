package handlers

import (
	"fmt"
	"intcode-discord-bot/arguments"
	"intcode-discord-bot/commands"
	"log"
	"strings"

	dgo "github.com/bwmarrin/discordgo"
)

// BOT_PREFIX is the prefix for any command sent to the bot
const BOT_PREFIX string = "!intcode"

// CreateMessage handles a message coming from discord
func CreateMessage(s *dgo.Session, m *dgo.MessageCreate) {

	m.Content = strings.TrimSpace(m.Content)
	// Check if the message is intended for this bot
	if !strings.HasPrefix(m.Content, BOT_PREFIX) {
		return
	}

	if m.Content == BOT_PREFIX {
		// User called the bot but didn't specify a command,
		// assume help command
		m.Content = BOT_PREFIX + " help"
	}

	// Skip the prefix
	m.Content = m.Content[len(BOT_PREFIX)+1:]

	// Process command and figure out the reply to send
	c := arguments.ParseCommandArguments(m.Content)

	var message string
	var cmdErr error

	switch c.Command {
	case "run":
		message, cmdErr = commands.Run(c)
	case "help":
		message = commands.Help(BOT_PREFIX)
	default:
		cmdErr = fmt.Errorf("command %s not recognized. Type `!f1 %s` for a full list of commands available", c.Command, BOT_PREFIX)
	}

	if cmdErr != nil {
		message = fmt.Sprintf("ERROR: %v", cmdErr)
	}

	// trim message to respect discord limits
	if len(message) > 2000 {
		message = message[:1950]
		message += " (...)"
	}

	// Send the message
	_, sendErr := s.ChannelMessageSend(m.ChannelID, message)
	if sendErr != nil {
		log.Printf("error sending message to discord: %v", sendErr)
	}

	log.Printf("Guild: %v | Author: %v(%v) | Command: %v | CmdErr: %v | SendErr: %v", m.GuildID, m.Author.ID, m.Author.Username, m.Content, cmdErr, sendErr)

}
