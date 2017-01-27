package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

/*
  .chess? maybe some other board game
  .clip
  .trivia
  image bot
  .delete
  channel stats
*/

func init() {
	rand.Seed(time.Now().UnixNano())
}

const tokenPath string = "TOKEN"
const prefix string = "."
const rollRange int = 100

func main() {
	b, err := ioutil.ReadFile(tokenPath)

	if err != nil {
		log.Print("Error reading token")
		return
	}

	var token = string(bytes.TrimSpace(b))
	/*
		db, err := sql.Open("sqlite3", "./image.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		createStmt := `create table image (id integer not null primary key, submitter text, filepath text);`

		_, err = db.Exec(createStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, createStmt)
			return
		}
	*/
	// Create new session using bot token
	discord, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Print("Error creating Discord session", err)
		return
	}

	discord.AddHandler(ready)

	discord.AddHandler(messageReceived)

	discord.AddHandler(newServer)

	err = discord.Open()
	if err != nil {
		log.Print("Error opening Discord session", err)
	}

	fmt.Println("Bot operational.")
	<-make(chan struct{})
	return
}

// Beep Boop.
func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set the current game
	_ = s.UpdateStatus(0, "Beep Boop")
}

// Parses any time a new message is created
func messageReceived(s *discordgo.Session, m *discordgo.MessageCreate) {
	// username := m.Author.Username
	// channel := m.ChannelID

	if strings.HasPrefix(m.Content, prefix+"beep") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Boop")
	} else if strings.HasPrefix(m.Content, prefix+"close") {
		_ = s.Close()
		fmt.Println("Bot exiting")
		os.Exit(0)
	} else if strings.HasPrefix(m.Content, prefix+"help") {
		_, _ = s.ChannelMessageSend(m.ChannelID, helpList())
	} else if strings.HasPrefix(m.Content, prefix+"roll") {
		_, _ = s.ChannelMessageSend(m.ChannelID, strconv.Itoa(rand.Int()%rollRange))
	}
}

// Called whenever a new server is joined.
func newServer(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "DRAWING BALANCE")
			return
		}
	}
}

// Display a list of commands.
func helpList() string {
	b, err := ioutil.ReadFile("help file")
	if err != nil {
		log.Print("Error reading the help file")
		return ""
	}
	var help = string(bytes.TrimSpace(b))
	return help
}
