package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func SendQuote(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	if rand.Intn(1000) < 1000 {
		unusedReplies := config.UnusedQuotes
		usedReplies := config.UsedQuotes

		var message string

		message, unusedReplies, usedReplies = reply.GetRandomMessage(unusedReplies, usedReplies)

		config.UnusedQuotes = unusedReplies
		config.UsedQuotes = usedReplies

		session.ChannelMessageSend(event.ChannelID, message)
	}
}
