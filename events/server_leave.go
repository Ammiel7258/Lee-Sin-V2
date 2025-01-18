package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"
	"os"

	"github.com/bwmarrin/discordgo"
)

func ServerLeave(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	channelID := os.Getenv("LOBBY_ID")
	unusedReplies := config.UnusedMemberLeaveMessages
	usedReplies := config.UsedMemberLeaveMessages

	var message string

	message, unusedReplies, usedReplies = reply.GetRandomMessage(unusedReplies, usedReplies)

	config.UnusedMemberLeaveMessages = unusedReplies
	config.UsedMemberLeaveMessages = usedReplies

	session.ChannelMessageSend(channelID, message)
}
