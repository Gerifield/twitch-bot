package bot

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type IdeaSaver interface {
	Save(id string, text string) error
}

// Bot .
type Bot struct {
	ideaSaver IdeaSaver
}

// New .
func New(ideaSaver IdeaSaver) *Bot {
	return &Bot{
		ideaSaver: ideaSaver,
	}
}

// Handler .
func (b *Bot) Handler(command string, params []string) error {

	switch command {
	case "!ötlet":
		return b.handleIdea(strings.Join(params[:], " "))
	}

	return nil
}

func (b *Bot) handleIdea(idea string) error {
	fmt.Println("Otlet:", idea)
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	return b.ideaSaver.Save(id.String(), idea)
}
