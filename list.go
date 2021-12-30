package main

import (
	"encoding/json"
	"os"
	"strings"
)

var users []string = nil

func init() {
	bytes, err := os.ReadFile("users.json")
	must(err)
	must(json.Unmarshal(bytes, &users))
}

func addUsertoList(username string) error {
	username = strings.ToLower(username)
	users = append(users, username)
	return saveUserlist()
}

func saveUserlist() error {
	bytes, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile("users.json", bytes, 0644)
}

func removeUserFromList(username string) error {
	username = strings.ToLower(username)
	tmp := []string{}
	for _, v := range users {
		if v != username {
			tmp = append(tmp, v)
		}
	}
	users = tmp[:]
	return saveUserlist()
}

func userIsOnList(username string) bool {
	username = strings.ToLower(username)
	for _, v := range users {
		if v == username {
			return true
		}
	}
	return false
}
