package quotes

import (
	"math/rand"
)

func GetRandomMessage() string {
	messages := []string{
    `"Attend to Buisness today. Leave that street-side flower alone." -Great Wall fortune cookie`,
    `"Drinking 12 budweisers and playing league at 21 is at the same level" -Kamden`,
    `"Use the dome before the phone" -BigQ`,
    `"no bill cosby shit but if niggas is sleepin then fuck em" -Karl`,
    `"bruh where this hoe irma at its still so quiet" -Karl`,
    `"15 YEAR OLD GIRLS ARE THE BEST" -Kamden`,
    `"squad up?" -Teon`,
    `"Is the calculator in degrees?" -Weg`,
    `"Fact of the Day: You niggas is gay" -Word of the Day (bot)`,
    `"I like my bread burnt" -Lady at subway`,
    `"Got to poop" -Justin`,
    `"yagga" -Grant (his first ever discord message)`,
    `"Life is too short to work at FedEx" -Hubie`,
    `"Been tried to warn u bout these niggas but u was too busy fucking em" -Karl`,
    `"Bitches in the army worried more about an IUD than an IED" -Seidon at 12:52 EST`,
    `"When you in college, age dont matter" -Grant`,
    `"He so ass" -Justin (when asked about Trae Young)`,
    `"My dumbass ate a peanut butter and jelly sandwich and got peanut butter in my wisdom tooth hole" -Justin (on his biggest regret)`,
    `"yess Karl my body sore lf. Fucking Ricky big ass have my back fuck up" -Kamden`,
    `"i lost vs aidan last week man anything possible" -Justin`,
    `"straight back to the chip zone ig 😓" -Axel (Saka got injured)`,
    `"SLUT ME OUTTTT" -Khalil (unprompted)`,
	}

	index := rand.Intn(len(messages))

	return messages[index]
}
