package main

import "time"

// track last time a user used a command, and the
// last time a channel had a command used. Then
// just see how much time has passed since then
// to see if they are within the limit

var (
	userLastuse         map[string]time.Time = make(map[string]time.Time)
	channelLastuse      map[string]time.Time = make(map[string]time.Time)
	perUserRateLimit                         = time.Second * 12
	perChannelRateLimit                      = time.Second * 5
)

func UserCheckLimit(username string) bool {
	return time.Since(userLastuse[username]) > perUserRateLimit
}
func UserRegisterLimit(username string) {
	userLastuse[username] = time.Now()
}
func ChannelCheckLimit(channel string) bool {
	return time.Since(channelLastuse[channel]) > perChannelRateLimit
}
func ChannelRegisterLimit(channel string) {
	channelLastuse[channel] = time.Now()
}
func RegisterUserChannelComboAllin1(username, channel string) {
	UserRegisterLimit(username)
	ChannelRegisterLimit(channel)
}
