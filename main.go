package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
)

func main() {
	var botClient *twitch.Client
	{ // scoped so variables containing private information arent even availible outside
		_, _, otheruser, otherpass, err := getLogin()
		must(err)
		botClient = twitch.NewClient(otheruser, otherpass)
	}

	// do stuff
	botClient.Join("turtoise", "quinndt", "snappingbot", "pajlada")
	announceSnappingbotGone := func(message twitch.PrivateMessage) {
		// fmt.Println("anounce")
		command := strings.Split(message.Message, " ")[0]
		if strings.HasPrefix(message.Message, "+") &&
			command != "+" &&
			command != "+ratio" {
			// fmt.Println("anounce2")
			RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
			botClient.Say(message.Channel, "snappingbot is offline sorry for any inconvenience - @t󠀀urtoise")
		}
	}
	reactToPajbot := func(message twitch.PrivateMessage) {
		if message.User.Name == "pajbot" && message.Action && message.Message == "pajaS 🚨 ALERT" {
			botClient.Say(message.Channel, "/me pajaVanish 🚨 ALERT RECEIVED")
		}
	}
	botClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Printf("[#%s] <%s>: %s\n", message.Channel, message.User.Name, message.Message)
		if !UserCheckLimit(message.User.Name) || !ChannelCheckLimit(message.Channel) {
			fmt.Println("skip because ratelimit")
			return
		}
		if message.Channel == "turtoise" || message.Channel == "snappingbot" || message.Channel == "quinndt" {
			announceSnappingbotGone(message)
		} else if message.Channel == "pajlada" {
			reactToPajbot(message)
		}
	})
	fruit := []string{"🍇", "🍈", "🍉", "🍊", "🍋", "🍌", "🍍", "🥭", "🍎", "🍏", "🍐", "🍑", "🍒", "🍓", "🥝"}
	rand.Seed(time.Now().Unix())
	botClient.OnConnect(func() {
		fmt.Println("Connected bot")
		botClient.Say("turtoise", "/me turtMunch "+fruit[rand.Intn(len(fruit))])
	})
	err := botClient.Connect()
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func getLogin() (user string, pass string, otheruser string, otherpass string, err error) {
	_user, err := ioutil.ReadFile("username")
	if err != nil {
		return
	}
	_pass, err := ioutil.ReadFile("password")
	if err != nil {
		return
	}

	_otheruser, err := ioutil.ReadFile("otherusername")
	if err != nil {
		return
	}

	_otherpass, err := ioutil.ReadFile("otherpassword")
	if err != nil {
		return
	}

	return string(_user), string(_pass), string(_otheruser), string(_otherpass), nil
}
