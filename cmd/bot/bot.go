package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gerifield/twitch-bot/bot"
	"github.com/gerifield/twitch-bot/command/jatek"
	"github.com/gerifield/twitch-bot/model"
	"github.com/gerifield/twitch-bot/twitch"

	"gopkg.in/irc.v3"
)

func regCommands(b *bot.Bot) {
	// b.Register("!vod", vods.Handle)
	// b.Register("!kappa", kappa.Handle)
	b.Register("!jatek", jatek.Handle)
}

func main() {
	channelName := flag.String("channel", "gerifield", "Twitch channel name")
	botName := flag.String("botName", "CoderBot42", "Bot name")
	clientID := flag.String("clientID", "", "Twitch App ClientID")
	clientSecret := flag.String("clientSecret", "", "Twitch App clientSecret")
	flag.Parse()

	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Println(err)
		return
	}

	tl := twitch.New(*clientID, *clientSecret)
	token, err := tl.GetToken()
	if err != nil {
		log.Println(err)
		return
	}

	myBot := bot.New()
	regCommands(myBot)

	config := irc.ClientConfig{
		Nick: *botName,
		Pass: "oauth:" + token.AccessToken,
		User: *botName,
		Name: *botName,
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			log.Println("incoming message", m)

			if m.Command == "001" {
				// 001 is a welcome event

				// We JOIN the given channel
				_ = c.Write("JOIN #" + *channelName)
			} else if m.Command == "PING" { // Handle occasional PINGs
				fmt.Println("PING HANDLED", c.Write("PONG :tmi.twitch.tv"))

			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				msg := model.ParseMessage(m)

				if !strings.HasPrefix(msg.Command(), "!") {
					return
				}

				resp, err := myBot.Handler(msg)
				if err != nil {
					log.Println(err)
				}

				if resp != "" {
					err = c.WriteMessage(&irc.Message{
						Command: "PRIVMSG",
						Params: []string{
							m.Params[0],
							resp,
						},
					})
					if err != nil {
						log.Println(err)
					}
				}
			}
		}),
	}

	// Create the client
	client := irc.NewClient(conn, config)
	client.CapRequest("twitch.tv/tags", true)     // Ask for tags
	client.CapRequest("twitch.tv/commands", true) // Ask for Twitch specific commands
	err = client.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
