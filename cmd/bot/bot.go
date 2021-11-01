package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gerifield/twitch-bot/bot"
	"github.com/gerifield/twitch-bot/db"

	bolt "go.etcd.io/bbolt"
	"gopkg.in/irc.v3"
)

func main() {
	channelName := flag.String("channel", "gerifield", "Twitch channel name")
	botName := flag.String("botName", "Suba", "Bot name")
	token := flag.String("token", "", "Twitch oauth token")
	flag.Parse()

	b, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() { _ = b.Close() }()

	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Println(err)
		return
	}

	myBot := bot.New(db.New(b))

	config := irc.ClientConfig{
		Nick: *botName,
		Pass: *token,
		User: *botName,
		Name: *botName,
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			fmt.Println(m)
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				_ = c.Write("JOIN #" + *channelName)
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				msgs := strings.Split(m.Trailing(), " ")
				if len(msgs) < 2 {
					return
				}

				if !strings.HasPrefix(msgs[0], "!") {
					return
				}

				err = myBot.Handler(msgs[0], msgs[1:])
				if err != nil {
					log.Println(err)
					return
				}
			}
		}),
	}

	// Create the client
	client := irc.NewClient(conn, config)
	err = client.Run()
	if err != nil {
		log.Println(err)
		return
	}

}
