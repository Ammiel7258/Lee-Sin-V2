package handlemessage

import (
	"os"
	"strings"
)

func SendFeatureMenu(messageData MessageData) {
	s, mc := messageData.session, messageData.messageCreate
	var sb strings.Builder

	if os.Getenv("KARL_ID") == mc.Author.ID {
		s.ChannelMessageSendReply(mc.ChannelID, "🖕🖕🖕 figure it out yourself https://github.com/Ammiel7258/Lee-Sin-V2 🖕🖕🖕", mc.Reference())
		return
	}

	sb.WriteString("\nHere is a list of leesin's features: \n")
	sb.WriteString("```")
	sb.WriteString(`
    • send funny discord quotes to the text chat on occassion
    • send happy birthday messages (wip)
    • notify users when a member joins or leaves the discord server
    • notify users when a message is edited or deleted`)
	sb.WriteString("```")

	featuresMessage := sb.String()

	s.ChannelMessageSend(mc.ChannelID, featuresMessage)
}
