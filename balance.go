package balancebot

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var token string

const prefix string = "!"

func main() {
	if token == "" {
		fmt.Println("No token provided. Usage: mechinabot -t <bot token>")
		return
	}
	// Create new session using bot token
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Print("Error creating Discord session", err)
		return
	}

	discord.AddHandler(ready)

	discord.AddHandler(messageCreate)

	discord.AddHandler(guildCreate)

	err = discord.Open()
	if err != nil {
		log.Print("Error opening Discord session", err)
	}

	fmt.Println("Bot operational.")
	<-make(chan struct{})
	return
}

// Initialize the bot.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the current game
	_ = s.UpdateStatus(0, "Beep Boop")
}

// Parses any time a new message is created on any channel that
// the bot is authenticated for.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// username := m.Author.Username
	// channel := m.ChannelID

	if strings.HasPrefix(m.Content, prefix+"Beep") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Boop")
	}
}

// Called whenever a new guild is joined.
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Bot is ready")

			return
		}
	}
}
