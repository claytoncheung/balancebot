package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func parse(s *discordgo.Session, m *discordgo.MessageCreate) {
	re := regexp.MustCompile("o(\\s*)3o")

	switch {
	case strings.HasPrefix(m.Content, prefix+"beep"):
		_, _ = s.ChannelMessageSend(m.ChannelID, "Boop")
	case strings.HasPrefix(m.Content, prefix+"close"):
		_ = s.Close()
		fmt.Println("Bot exiting")
		os.Exit(0)
	case strings.HasPrefix(m.Content, prefix+"help"):
		_, _ = s.ChannelMessageSend(m.ChannelID, helpList())
	case strings.HasPrefix(m.Content, prefix+"roll"):
		_, _ = s.ChannelMessageSend(m.ChannelID, strconv.Itoa(rand.Int()%rollRange))
	default:

	}
	if re.MatchString(m.Content) {
		_, _ = s.ChannelMessageEdit(m.ChannelID, m.ID, re.ReplaceAllString(m.Content, "frick u"))
	}

}
