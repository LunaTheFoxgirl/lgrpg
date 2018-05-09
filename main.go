package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type BotConfiguration struct {
	BotName  string `json:"name"`
	BotToken string `json:"token"`
}

var counter int = 0

func OnMessageCreate(sess *discordgo.Session, event *discordgo.MessageCreate) {
	// Ignore messages sent from the bot itself.
	if event.Author.ID == sess.State.User.ID {
		return
	}

	fmt.Println(event.Message.Content)

	// Example reaction
	if event.Message.Content == "Ping" {
		msg, err := sess.ChannelMessageSend(event.ChannelID, "Hey <@"+event.Author.ID+">, pong!")
		if err != nil {
			fmt.Println("Failure! ", err.Error())
			return
		}
		fmt.Println(msg.ChannelID)
	}
	if event.Message.Content == "New Name" {
		msg, err := sess.ChannelMessageSend(event.ChannelID, "Hey <@"+event.Author.ID+">, name is: "+GenerateRandomSpacename(counter))
		if err != nil {
			fmt.Println("Failure! ", err.Error())
			return
		}
		fmt.Println(msg.ChannelID)
		counter++;
	}
}

func OnReady(sess *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Ready to handle the game!")
}

func main() {
	fmt.Println("Starting client...")
	// prep vars
	var config BotConfiguration = BotConfiguration{}

	// Read configuration file
	dat, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("File read error (config.json) ", err.Error())
		return
	}

	// Unmarshal text file
	err = json.Unmarshal(dat, &config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Start discord session
	session, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Add handlers
	session.AddHandler(OnReady)
	session.AddHandler(OnMessageCreate)

	// Open connnection
	err = session.Open()
	if err != nil {
		fmt.Println("Could not establish connection with discord! ", err.Error())
		return
	}
	waitTillDeath()
	session.Close()
}

func waitTillDeath() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
