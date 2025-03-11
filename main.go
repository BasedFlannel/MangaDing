package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {

	authToken := readAuthFile("bot.key")

	sess, err := discordgo.New("Bot " + authToken)
	errorCheck(err)

	fmt.Println("Session initialized")

	sess.AddHandler(helloMesages)

	//18432 is Send Messages, read message history, and Embed Links
	sess.Identify.Intents = discordgo.IntentsGuildMessages

	err = sess.Open()
	errorCheck(err)

	fmt.Println("The bot is listening")

	//code to hold the thread until interrupt is sent
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func helloMesages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID, "world!")
	}
}

func readAuthFile(filename string) string {
	data, err := os.ReadFile(filename)
	errorCheck(err)
	return (string(data))
}

// function for checking errors, logs and panics if one comes up.
func errorCheck(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		panic(e)
	}
}
