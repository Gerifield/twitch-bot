package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"gopkg.in/irc.v3"
)

func main() {
	channelName := flag.String("channel", "gerifield", "Twitch channel name")
	botName := flag.String("botName", "Suba", "Bot name")
	token := flag.String("token", "", "Twitch oauth token")
	flag.Parse()

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() { _ = db.Close() }()

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
				_ = c.Write("JOIN #" + *channelName)
			} else if m.Command == "PRIVMSG" && c.FromChannel(m) {
				msgs := strings.Split(m.Trailing(), " ")
				if len(msgs) < 2 {
					return
				}

				if strings.ToLower(msgs[0]) == "!ötlet" {
					// Save idea
					idea := strings.Join(msgs[1:], " ")
					fmt.Println("Otlet:", idea)

					err = db.Update(func(tx *bolt.Tx) error {
						b, err := tx.CreateBucketIfNotExists([]byte("ideas"))
						if err != nil {
							return err
						}

						id, err := uuid.NewUUID()
						if err != nil {
							return err
						}

						log.Printf("Save idea with ID %s, %s", id.String(), idea)
						err = b.Put([]byte(id.String()), []byte(idea))
						return err
					})
					if err != nil {
						log.Println(err)
						return
					}
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
