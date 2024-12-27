package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"leesin-v2/features/handlemessages"

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

	/*
	  i want some kind of menu, maybe a switch statement?

	  AUTOMATED:
	  word of the day
	  sending happy birthday wishes
	  joining / leaving the server
	  editing / deleting messages

	  ON COMMAND:
	  !leesin help
	    -this will send the user a list of commands that bot can do
	  !leesin addBDay (mm/dd/yyyy)
	  !leesin removeBDay
	  !leesin updateBDay (mm/dd/yyyy)
	*/
  session.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
    handlemessages.HandleMessage(session, message)
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

/*
  users, err := features.ReadBirthdayData()
  if err != nil {
   log.Fatal(err)
  }
  features.PrintBirthdayData(users)
*/
