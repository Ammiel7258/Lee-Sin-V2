package handlemessage

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"leesin-v2/config"
	"leesin-v2/features/birthday"
	"leesin-v2/features/reply"

	"github.com/bwmarrin/discordgo"
)

type MessageData struct {
	session       *discordgo.Session
	messageCreate *discordgo.MessageCreate
}

const prefix string = config.Prefix

var messageData MessageData

func HandleMessages(s *discordgo.Session, mc *discordgo.MessageCreate) {
	messageData.session = s
	messageData.messageCreate = mc

	if mc.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(strings.ToLower(mc.Content), " ")

	// this should never trigger !
	if len(args) < 1 {
		s.ChannelMessageSend(mc.ChannelID, "🤯🤯🤯 you caused an error with lee sin v2 🤯🤯🤯")
		return
	}

	if args[0] == prefix {
		commandHelper(args[1:])
	} else {
		generalMessageHelper()
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
	if len(args) == 0 {
		invalidCommandMenu(args)
		return
	}

	command := args[0]

	switch command {
	case "help":
    helpMenu(args[1:])
		break
	case "features":
		featureMenu()
		break
	case "birthday":
		birthdayMenu(args)
		break
	default:
    fmt.Println("in default")
		invalidCommandMenu(args)
	}
}

func birthdayMenu(args []string) {

}

func featureMenu() {
	s, mc := messageData.session, messageData.messageCreate

	s.ChannelMessageSend(mc.ChannelID, "This is the features command")
  // 🖕
}

func helpMenu(args []string) {
	var sb strings.Builder

	s, mc := messageData.session, messageData.messageCreate

  if len(args) == 0 {
    invalidCommandMenu(nil)
    return
  }

  helpMessage := sb.String()
  s.ChannelMessageSend(mc.ChannelID, helpMessage)
}

func invalidCommandMenu(args []string) {
	var sb strings.Builder

	s, mc := messageData.session, messageData.messageCreate
  var command string = ""
  if (len(args) > 0) {
	  command = args[0]
  }

	if command == "" {
		sb.WriteString("This command requires arguments to work.")
	} else {
		sb.WriteString(fmt.Sprintf(`"%s" is not a valid command!`, command))
	}

	sb.WriteString("\nHere are the commands available to use after !leesin: \n")
	sb.WriteString("```")
	sb.WriteString(`
    1: !leesin help {command}\n
    2: !leesin features\n
    3: !leesin birthday {args}\n
    4: !leesin test
  `)
	sb.WriteString("```")

	helpMessage := sb.String()
	s.ChannelMessageSend(mc.ChannelID, helpMessage)
}

func generalMessageHelper() {
  var quote string
	s, mc := messageData.session, messageData.messageCreate

	if rand.Intn(1000) < 10 {
    quote, config.UnusedQuotes, config.UsedQuotes = reply.GetRandomMessage(config.UnusedQuotes, config.UsedQuotes)
		s.ChannelMessageSend(mc.ChannelID, quote)
	}

  fmt.Println(len(config.UsedQuotes), len(config.UnusedQuotes)) 
	// check if original timestamp (message id) is different from last edited timestamp
}

func handleBirthday(id string, date string) {
	s, mc := messageData.session, messageData.messageCreate
	channelID := mc.ChannelID

	err := validateBirthday(date)
	if err != nil {
		s.ChannelMessageSend(channelID, err.Error())
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
	return errors.New(
		fmt.Sprintf("%s is not a valid birthdate, the date must be formatted in this style: ```mm/dd/yyyy```", birthdate),
	)
}
