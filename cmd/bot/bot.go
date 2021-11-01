package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"gopkg.in/irc.v3"
)

func main() {
	channelName := flag.String("channel", "gerifield", "Twitch channel name")
	botName := flag.String("botName", "Suba", "Bot name")
	token := flag.String("token", "", "Twitch oauth token")
	flag.Parse()

	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Println(err)
		return
	}

	config := irc.ClientConfig{
		Nick: *botName,
		Pass: *token,
		User: *botName,
		Name: *botName,
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			fmt.Println(m)
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				c.Write("JOIN #" + *channelName)
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {

				fmt.Println(m)
				fmt.Println(m.Params)
				fmt.Println("Msg:", m.Trailing())
				fmt.Println(m.String())

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
