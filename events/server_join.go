package events

import (
	"leesin-v2/config"
	"leesin-v2/features/reply"
	"os"

	"github.com/bwmarrin/discordgo"
)

func ServerJoin(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	channelID := os.Getenv("LOBBY_ID")
	unusedReplies := config.UnusedMemberJoinMessages
	usedReplies := config.UsedMemberJoinMessages

	var message string

	message, unusedReplies, usedReplies = reply.GetRandomMessage(unusedReplies, usedReplies)

	config.UnusedMemberJoinMessages = unusedReplies
	config.UsedMemberJoinMessages = usedReplies

	session.ChannelMessageSend(channelID, message)
}
