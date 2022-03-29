package bot

import (
	"errors"

	"github.com/gerifield/twitch-bot/model"
)

// Handler .
type Handler func(*model.Message) (string, error)

// Bot .
type Bot struct {
	commands map[string]Handler
}

var (
	ErrNotFound = errors.New("command not found")
)

// New .
func New() *Bot {
	return &Bot{
		commands: make(map[string]Handler),
	}
}

func (b *Bot) Register(command string, handler Handler) {
	b.commands[command] = handler
}

// Handler .
func (b *Bot) Handler(msg *model.Message) (string, error) {
	h, ok := b.commands[msg.Command()]
	if !ok {
		return "", ErrNotFound
	}

	return h(msg)
}
