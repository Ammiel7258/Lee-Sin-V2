package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"

	"github.com/bwmarrin/discordgo"
)

func MessageEdit(session *discordgo.Session, event *discordgo.MessageUpdate) {
  var message string
  message, config.UnusedEditedMessageReplies, config.UsedEditMessageReplies = reply.GetRandomMessage(config.UnusedEditedMessageReplies, config.UsedEditMessageReplies)
	session.ChannelMessageSend(event.ChannelID, message)
}
