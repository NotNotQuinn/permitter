package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

func main() {
	user, pass, err := getLogin()
	must(err)
	client := twitch.NewClient(user, pass)

	// do stuff
	client.Join("michaelreeves")
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// fmt.Printf("[#%s] <%s>: %s\n", message.Channel, message.User.Name, message.Message)
		args := strings.Split(message.Message, " ")
		if message.User.Badges["moderator"] == 1 {
			// mod zone
			if args[0] == "!permitremove" {
				// remove user
				if len(args) < 2 {
					client.Say(message.Channel, "remove who")
					return
				}
				err := removeUserFromList(args[1])
				if err != nil {
					client.Say(message.Channel, "err: "+err.Error())
					return
				}
				client.Say(message.Channel, args[1]+" can no longer use !permitme Okayge")
				return
			}
			if args[0] == "!permitadd" {
				// add user
				if len(args) < 2 {
					client.Say(message.Channel, "add who")
					return
				}
				err := addUsertoList(args[1])
				if err != nil {
					client.Say(message.Channel, "err: "+err.Error())
					return
				}
				client.Say(message.Channel, args[1]+" can now use !permitme Okayge")
				return
			}
			if args[0] == "!permithelp" {
				// show help
				client.Say(message.Channel, "Mods can use !permitadd and !permitremove to manage who can use !permitme, !permitme posts !permit <user>.")
				return
			}
			return
		}
		if args[0] == "!permitme" {
			if userIsOnList(message.User.Name) {
				go func() {
					// send !permit <user>
					// time.Sleep(time.Duration((rand.Intn(4) + 1)) * time.Second) // wait 1-5 seconds, as recommended by dim
					client.Say(message.Channel, fmt.Sprintf("!permit %s", message.User.Name))
				}()
				return
			}
		}
	})

	client.OnConnect(func() {
		fmt.Println("Connected")
	})

	err = client.Connect()
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func getLogin() (user string, pass string, err error) {
	_user, err := ioutil.ReadFile("username")
	if err != nil {
		return
	}
	_pass, err := ioutil.ReadFile("password")
	if err != nil {
		return
	}
	return string(_user), string(_pass), nil
}
