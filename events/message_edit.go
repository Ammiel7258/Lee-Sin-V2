package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageEdit(session *discordgo.Session, event *discordgo.MessageUpdate) {
	unusedReplies := config.UnusedEditMessageReplies
	usedReplies := config.UsedEditMessageReplies

	var message string

	if containsALink(event.Content) {
		return
	}

	message, unusedReplies, usedReplies = reply.GetRandomMessage(unusedReplies, usedReplies)

	config.UnusedEditMessageReplies = unusedReplies
	config.UsedEditMessageReplies = usedReplies

	session.ChannelMessageSend(event.ChannelID, message)
}

func containsALink(message string) bool {
	linkRegex := `https?://[^\s]+|www\.[^\s]+`
	re := regexp.MustCompile(linkRegex)

	messageList := strings.Split(message, " ")

	for _, word := range messageList {
		isLink := re.MatchString(word)
		if isLink {
			return true
		}
	}

	return false
}
