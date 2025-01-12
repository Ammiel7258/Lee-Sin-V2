package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SetStatus(session *discordgo.Session, event *discordgo.Ready) {
  status := "😏🤯🕺😭"
  session.UpdateCustomStatus(status)
  fmt.Printf("Status set to: %s\n", status)
}
