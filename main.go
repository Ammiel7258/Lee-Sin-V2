package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"leesin-v2/events"
	"leesin-v2/features/handlemessage"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

  /* -----------------------figure out if there is a better way to handle a bunch of events----------------------- */
	session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		handlemessage.HandleMessages(session, message) // this needs to be redone...
	})

	session.AddHandler(func(session *discordgo.Session, status *discordgo.Ready) {
		events.SetStatus(session, status)
	})

	session.AddHandler(func(session *discordgo.Session, deleteEvent *discordgo.MessageDelete) {
		events.MessageDelete(session, deleteEvent)
	})

	session.AddHandler(func(session *discordgo.Session, editEvent *discordgo.MessageUpdate) {
		events.MessageEdit(session, editEvent)
	})

  // needs to be tested
	session.AddHandler(func(session *discordgo.Session, userJoinEvent *discordgo.GuildMemberAdd) {
		events.ServerJoin(session, userJoinEvent)
	})
  // needs to be tested
	session.AddHandler(func(session *discordgo.Session, userLeaveEvent *discordgo.GuildMemberRemove) {
		events.ServerLeave(session, userLeaveEvent)
	})


	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	fmt.Println("lee sin v2 activated! press CTRL+C to exit")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}
