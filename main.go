package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
)

func main() {
	var selfClient, botClient *twitch.Client
	{ // scoped so variables containing private information arent even availible outside
		user, pass, otheruser, otherpass, err := getLogin()
		must(err)
		selfClient = twitch.NewClient(user, pass)
		botClient = twitch.NewClient(otheruser, otherpass)
	}

	// do stuff
	selfClient.Join("michaelreeves", "turtoise", "quinndt", "snappingbot")
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
	permitUsers := func(message twitch.PrivateMessage) {
		args := strings.Split(message.Message, " ")
		if len(args) > 1 {
			args[1] = strings.ToLower(args[1])
		}
		if message.User.Badges["moderator"] == 1 {
			// mod zone
			if args[0] == "!permitremove" {
				// remove user
				if len(args) < 2 {
					RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
					selfClient.Say(message.Channel, "remove who")
					return
				}
				err := removeUserFromList(args[1])
				if err != nil {
					RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
					selfClient.Say(message.Channel, "err: "+err.Error())
					return
				}
				RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
				selfClient.Say(message.Channel, args[1]+" can no longer use !permitme Okayge")
				return
			}
			if args[0] == "!permitadd" {
				// add user
				if len(args) < 2 {
					RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
					selfClient.Say(message.Channel, "add who")
					return
				}
				err := addUsertoList(args[1])
				if err != nil {
					RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
					selfClient.Say(message.Channel, "err: "+err.Error())
					return
				}
				RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
				selfClient.Say(message.Channel, args[1]+" can now use !permitme Okayge")
				return
			}
			if args[0] == "!permithelp" {
				// show help
				RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
				selfClient.Say(message.Channel, "Mods can use !permitadd and !permitremove to manage who can use !permitme, !permitme posts !permit <user>.")
				return
			}
			return
		}
		if args[0] == "!permitme" {
			if userIsOnList(message.User.Name) {
				go func() {
					// send !permit <user>
					// time.Sleep(time.Duration((rand.Intn(4) + 1)) * time.Second) // wait 1-5 seconds, as recommended by dim
					RegisterUserChannelComboAllin1(message.User.ID, message.Channel)
					selfClient.Say(message.Channel, fmt.Sprintf("!permit %s", message.User.Name))
				}()
				return
			}
		}
	}
	selfClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Printf("[#%s] <%s>: %s\n", message.Channel, message.User.Name, message.Message)
		if !UserCheckLimit(message.User.Name) || !ChannelCheckLimit(message.Channel) {
			fmt.Println("skip because ratelimit")
			return
		}
		if message.Channel == "michaelreeves" {
			permitUsers(message)
		} else if message.Channel == "turtoise" || message.Channel == "snappingbot" || message.Channel == "quinndt" {
			announceSnappingbotGone(message)
		}
	})
	selfClient.OnConnect(func() {
		fmt.Println("Connected self")
	})

	botClient.OnConnect(func() {
		fmt.Println("Connected bot")
	})
	go func() {
		err := botClient.Connect()
		must(err)
	}()
	err := selfClient.Connect()
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
