package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

// Bot .
type Bot struct {
	db *bolt.DB
}

// New .
func New(db *bolt.DB) *Bot {
	return &Bot{
		db: db,
	}
}

// Handler .
func (b *Bot) Handler(command string, params []string) error {

	switch command {
	case "!ötlet":
		b.handleIdea(strings.Join(params[:], " "))
	}

	return nil
}

func (b *Bot) handleIdea(idea string) {
	fmt.Println("Otlet:", idea)

	err := b.db.Update(func(tx *bolt.Tx) error {
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
