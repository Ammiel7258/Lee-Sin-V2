package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"

	"github.com/bwmarrin/discordgo"
)

func MessageDelete(session *discordgo.Session, event *discordgo.MessageDelete) {
  var message string
  message, config.UnusedDeleteMessageReplies, config.UsedDeleteMessageReplies = reply.GetRandomMessage(config.UnusedDeleteMessageReplies, config.UsedDeleteMessageReplies)
	session.ChannelMessageSend(event.ChannelID, message)
}
