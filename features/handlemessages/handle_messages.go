package handlemessages

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"leesin-v2/config"
	"leesin-v2/features/birthday"
	"leesin-v2/features/quotes"

	"github.com/bwmarrin/discordgo"
)

type MessageData struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
}

const prefix string = config.Prefix

func HandleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	isCommandMessage := false

	args := strings.Split(strings.ToLower(message.Content), " ")

  if len(args) < 1 {
    session.ChannelMessageSend(message.ChannelID, "🤯🤯🤯 you found an error in lee sin v2 🤯🤯🤯")
    return
  }

	if args[0] == prefix {
		isCommandMessage = true
	}

	if isCommandMessage {
    commandHelper(args[1:])
	} else {
		generalMessageHelper(session, message, message.ChannelID)
	}

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

func commandHelper(args []string) {
  // lets work on the birthday commands
  n := len(args)

  fmt.Println(n, args)


}

func generalMessageHelper(session *discordgo.Session, message *discordgo.MessageCreate, channelID string) {
	// on certain chance, send a quote
  randNum := rand.Intn(1000)

  // now we send a random message to the server
  if (randNum < 25) {
    message := quotes.GetRandomMessage()
    session.ChannelMessageSend(channelID, message)
  }

  // check if original timestamp (message id) is different from last edited timestamp

}

func handleBirthday(id string, date string, session *discordgo.Session, channelID string) {

	err := validateBirthday(date)
	if err != nil {
		session.ChannelMessageSend(channelID, err.Error())
		return
	}

	err = birthday.AddBirthday(id, "testval", date)
	if err != nil {
		// handle the error here
		return
	}
}

// validateBirthday is a function that will return an error if the birthdate is not a valid date, if the date is valid, return nil
func validateBirthday(birthdate string) error {
	return errors.New(fmt.Sprintf("%s is not a valid birthdate, the date must be formatted in this style: ```mm/dd/yyyy```", birthdate))
}

/* session.AddHandler(func(s *discordgo.Session, msg *discordgo.MessageCreate) {

		args := strings.Split(msg.Content, " ")

		p := strings.ToLower(args[0])
		command := strings.ToLower(args[1])

		if p == prefix {
			switch command {
			case "help":
				s.ChannelMessageSend(msg.ChannelID, "This is the help menu:")
			case "test":
				s.ChannelMessageSend(msg.ChannelID, "This is the test message")
			default:
        s.ChannelMessageSend(msg.ChannelID, "This is not a valid command. Use !leesin help to see a list of commands")
			}
		}

	}) */
/* if pre == prefix {
	switch command {
	case "help":
		session.ChannelMessageSend(message.ChannelID, "This is the help menu:")
	case "remove_birthday":
		if len(args) != 2 {
			session.ChannelMessageSend(message.ChannelID, "Incorrect message format! Use: ```!leesin remove_birthday {mm/dd/yyyy}```")
			break
		}
	case "edit_birthday":
		if len(args) != 3 {
			session.ChannelMessageSend(message.ChannelID, "Incorrect message format! Use: ```!leesin edit_birthday mm/dd/yyyy```")
			break
		}
	case "add_birthday":
		if len(args) != 3 {
			session.ChannelMessageSend(message.ChannelID, "Incorrect message format! Use: ```!leesin add_birthday mm/dd/yyyy```")
			break
		}
		date := args[2]
		handleBirthday(message.Author.ID, date, session, message.ChannelID)
	default:
		session.ChannelMessageSend(message.ChannelID, "This is not a valid command. Use !leesin help to see a list of commands")
	}
} */
