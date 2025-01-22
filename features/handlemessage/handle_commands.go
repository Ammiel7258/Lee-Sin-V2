package handlemessage

import (
	"fmt"
	"strings"

	"leesin-v2/config"

	"github.com/bwmarrin/discordgo"
)

type MessageData struct {
	session       *discordgo.Session
	messageCreate *discordgo.MessageCreate
}

const prefix string = config.Prefix

var messageData MessageData

func ParseMessage(s *discordgo.Session, mc *discordgo.MessageCreate) {
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
	}
}

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
		SendFeatureMenu(messageData)
		break
	case "birthday":
		birthdayMenu(args[1:])
		break
	default:
		fmt.Println("in default")
		invalidCommandMenu(args)
	}
}

func birthdayMenu(args []string) {
	if !birthdayCommandIsValid(args) {
		invalidCommandMenu(append(args[len(args):], "")) // clears the slice and appends ""
		return
	}

	switch args[0] {
	case "add":
		handleBirthday("add", args[1])
		break
	case "edit":
		handleBirthday("edit", args[1])
		// fmt.Println(fmt.Sprintf("editing %s", args[1]))
		break
	case "remove":
		// fmt.Println(fmt.Sprintf("removing birthday"))
		break
	default:
		invalidCommandMenu(nil)
		break
	}

}

func helpMenu(args []string) {
}

func invalidCommandMenu(args []string) {
	var sb strings.Builder

	s, mc := messageData.session, messageData.messageCreate
	var command string = ""
	if len(args) > 0 {
		command = args[0]
	}

	if command == "" {
		sb.WriteString("Invalid arguments!")
	} else {
		sb.WriteString(fmt.Sprintf(`"%s" is not a valid command!`, command))
	}

	sb.WriteString("\nHere are the commands available to use after !leesin: \n")
	sb.WriteString("```")
	sb.WriteString(`
    • !leesin help {command name}
    • !leesin features
    • !leesin birthday {command type}
    • !leesin test`)
	sb.WriteString("```")

	helpMessage := sb.String()
	s.ChannelMessageSend(mc.ChannelID, helpMessage)
}

func generalMessageHelper() {
}

func handleBirthday(id string, date string) {
	/*
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
	*/
}

// makes sure that the arguments are valid
func birthdayCommandIsValid(args []string) bool {
	numCommands := len(args)
	/*
	   remove
	   add date
	   edit date
	*/
	fmt.Println(args)
	if numCommands != 1 && numCommands != 2 {
		return false
	}

	// this commented code should be redundant with the switch statement
	/* if args[0] != "remove" && args[0] != "add" && args[0] != "edit" {
		return false
	} */

	switch args[0] {
	case "remove":
		if numCommands != 1 {
			return false
		}
		break
	case "edit":
		if numCommands != 2 {
			return false
		}
		break
	case "add":
		if numCommands != 2 {
			return false
		}
		if !validateBirthday(args[1]) {
			return false
		}
		break
	default:
		return false
	}

	return true
}

// validateBirthday is a function that will return an error if the birthdate is not a valid date, if the date is valid, return nil
func validateBirthday(birthdate string) bool {
	/* return errors.New(
		fmt.Sprintf("%s is not a valid birthdate, the date must be formatted in this style: ```mm/dd/yyyy```", birthdate),
	) */
	return true
}

/*
  i want some kind of menu, maybe a switch statement?

  AUTOMATED:
  word of the day
  sending happy birthday wishes

  ON COMMAND:
  !leesin help
    -this will send the user a list of commands that bot can do
  !leesin addBDay (mm/dd/yyyy)
  !leesin removeBDay
  !leesin updateBDay (mm/dd/yyyy)
*/
