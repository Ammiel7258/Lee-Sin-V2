package reply

import "math/rand"

/*
  takes two string slices, one that is empty, and the other that is full. this function will get a random message from the
  full slice and return it. it will also move the now used message into the empty slice. this function returns the string
  and the two slices after being altered
*/
func GetRandomMessage(unusedMessages []string, usedMessages []string) (string, []string, []string) {
	if len(unusedMessages) < 1 {
		unusedMessages = append(unusedMessages, usedMessages[0:]...)
		usedMessages = usedMessages[:0]
	}

	index := rand.Intn(len(unusedMessages))
	message := unusedMessages[index]

	usedMessages = append(usedMessages, message)
	unusedMessages = append(unusedMessages[:index], unusedMessages[index+1:]...)

	return message, unusedMessages, usedMessages
}
