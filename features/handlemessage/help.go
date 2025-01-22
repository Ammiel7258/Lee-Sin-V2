package handlemessage

import (
	"os"
	"strings"
)

func SendHelpMenu(messageData MessageData, args []string) {
	var sb strings.Builder
	karlID := os.Getenv("KARL_ID")

	s, mc := messageData.session, messageData.messageCreate

	if karlID == mc.Author.ID {
		s.ChannelMessageSendReply(mc.ChannelID, "🖕🖕🖕 figure it out yourself https://github.com/Ammiel7258/Lee-Sin-V2 🖕🖕🖕", mc.Reference())
		return
	}

	if len(args) == 0 {
		invalidCommandMenu(nil)
		return
	}

	helpMessage := sb.String()
	s.ChannelMessageSend(mc.ChannelID, helpMessage)
	s.ChannelMessageSend(mc.ChannelID, " ")
}
