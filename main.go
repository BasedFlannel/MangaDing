package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	restTest := restGet("https://api.restful-api.dev/objects")
	log.Println(restTest)
	//runBot()
}

func runBot() {

	//import a discord auth token from bot.key and instantiate a new bot with it
	authToken := loadFile("bot.key.fart")
	sess, err := discordgo.New("Bot " + authToken)
	errorCheck(err)

	//add message handler, intents, and open the session
	sess.AddHandler(helloMesages)
	sess.Identify.Intents = discordgo.IntentsGuildMessages
	err = sess.Open()
	errorCheck(err)
	log.Println("The bot is listening")

	//code to hold the thread until interrupt is sent
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// Primary message handler
func helloMesages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID, "world!")
	}
}

func loadFile(filename string) string {
	data, err := os.ReadFile(filename)
	errorCheck(err)
	return (string(data))
}

// basic error check function to throw panic when nothing more is required
func errorCheck(e error) {
	if e != nil {
		log.Panic(e)
	}
}
