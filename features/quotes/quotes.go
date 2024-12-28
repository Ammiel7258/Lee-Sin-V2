package quotes

import (
	"leesin-v2/config"
	"math/rand"
)

var usedMessages = []string{}
var unusedMessages = config.UnusedMessages

func GetRandomMessage() string {
	if len(unusedMessages) < 1 {
		unusedMessages = append(unusedMessages, usedMessages[0:]...)
		usedMessages = usedMessages[:0]
	}

	index := rand.Intn(len(unusedMessages))
	message := unusedMessages[index]

	usedMessages = append(usedMessages, message)
	unusedMessages = append(unusedMessages[:index], unusedMessages[index+1:]...)

	return message
}
