package events

import "github.com/bwmarrin/discordgo"

func SetStatus(session *discordgo.Session, event *discordgo.Ready) {
	status := "😏🤯🕺😭"
	session.UpdateCustomStatus(status)
}
