package bot

import "errors"

// Handler .
type Handler func([]string) (string, error)

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
func (b *Bot) Handler(command string, params []string) (string, error) {
	h, ok := b.commands[command]
	if !ok {
		return "", ErrNotFound
	}

	return h(params)
}
