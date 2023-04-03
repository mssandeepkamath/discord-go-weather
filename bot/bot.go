package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	OpenWeatherToken string
	BotToken         string
)

func Run() {

	//create a session
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	//add a event handler
	discord.AddHandler(newMessage)

	//open session

	discord.Open()
	defer discord.Close()

	//Run until the code terminater

	fmt.Println("Bot running....")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	//prevent bot responding to its own message

	if message.Author.ID == discord.State.User.ID {
		return
	}
	//respond to user message

	switch {
	case strings.Contains(message.Content, "help"):
		discord.ChannelMessageSend(message.ChannelID, "I can help you with weather!, Use !zip <6 digit zip code> (Currently supports India only)")
	case strings.Contains(message.Content, "!zip"):
		weatherData := getCurrentWeather(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, weatherData)
	}
}

func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error as occured")
	}
}
