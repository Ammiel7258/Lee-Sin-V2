package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"

	"github.com/bwmarrin/discordgo"
)

func MessageDelete(session *discordgo.Session, event *discordgo.MessageDelete) {
  unusedReplies := config.UnusedDeleteMessageReplies
  usedReplies := config.UsedDeleteMessageReplies

	var message string

	message, unusedReplies, usedReplies = reply.GetRandomMessage(unusedReplies, usedReplies)

  config.UnusedDeleteMessageReplies = unusedReplies
  config.UsedDeleteMessageReplies = usedReplies

	session.ChannelMessageSend(event.ChannelID, message)

}
