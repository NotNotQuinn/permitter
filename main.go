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
		otheruser, otherpass, err := getLogin()
		must(err)
		botClient = twitch.NewClient(otheruser, otherpass)
	}

	// do stuff
	botClient.Join("turtoise", "quinndt", "snappingbot", "pajlada", "supinic")
	// announceSnappingbotGone := func(message twitch.PrivateMessage) {
	// 	// fmt.Println("anounce")
	// 	command := strings.ToLower(strings.Split(message.Message, " ")[0])
	// 	if strings.HasPrefix(message.Message, "+") &&
	// 		command != "+" &&
	// 		command != "+l" &&
	// 		command != "+ratio" {
	// 		// fmt.Println("anounce2")
	// 		RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
	// 		botClient.Say(message.Channel, "snappingbot is offline sorry for any inconvenience - @t󠀀urtoise")
	// 	}
	// }
	reactToPajbot := func(message twitch.PrivateMessage) {
		if message.User.Name == "pajbot" && message.Action && message.Message == "pajaS 🚨 ALERT" {
			UserRegisterLimit("pajbot")
			botClient.Say(message.Channel, "/me pajaVanish 🚨 ALERT RECEIVED")
		}
		if message.User.Name == "mm_sutilitybot" && strings.StartsWith(message.Message, "/announce 🅱") {
			UserRegisterLimit("mm_sutilitybot")
			botClient.Say(message.Channel, "/ /announce 💿"
		}
	}
	reactToSupibot := func(message twitch.PrivateMessage) {
		if message.User.Name == "supibot" && message.Message == "ppCircle" {
			UserRegisterLimit("supibot")
			botClient.Say(message.Channel, "ppOverheat ￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼ ￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼￼  ppCircleHeat")
		}
	}
	botClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Printf("[#%s] <%s>: %s\n", message.Channel, message.User.Name, message.Message)
		if !UserCheckLimit(message.User.Name) || !ChannelCheckLimit(message.Channel) {
			fmt.Println("skip because ratelimit")
			return
		}
		if message.Channel == "turtoise" || message.Channel == "snappingbot" || message.Channel == "quinndt" {
			// qannounceSnappingbotGone(message)
		} else if message.Channel == "pajlada" {
			reactToPajbot(message)
		} else if message.Channel == "supinic" {
			reactToSupibot(message)
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

func getLogin() (otheruser string, otherpass string, err error) {
	_otheruser, err := ioutil.ReadFile("otherusername")
	if err != nil {
		return
	}

	_otherpass, err := ioutil.ReadFile("otherpassword")
	if err != nil {
		return
	}

	return string(_otheruser), string(_otherpass), nil
}
