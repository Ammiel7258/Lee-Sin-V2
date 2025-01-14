package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"

	"github.com/bwmarrin/discordgo"
)

func MessageEdit(session *discordgo.Session, event *discordgo.MessageUpdate) {
  unusedReplies := config.UnusedEditMessageReplies
  usedReplies := config.UsedEditMessageReplies

	var message string

	message, unusedReplies, usedReplies = reply.GetRandomMessage(unusedReplies, usedReplies)

  config.UnusedEditMessageReplies = unusedReplies
  config.UsedEditMessageReplies = usedReplies

	session.ChannelMessageSend(event.ChannelID, message)
}
