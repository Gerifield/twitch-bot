package model

import (
	"strings"

	"gopkg.in/irc.v3"
)

// Message wrapper for the raw IRC one
type Message struct {
	original *irc.Message
	parts    []string
}

// ParseMessage parse an IRC message
func ParseMessage(msg *irc.Message) *Message {
	return &Message{
		original: msg,
		parts:    strings.Split(msg.Trailing(), " "),
	}
}

// Command returns the parse message's command
func (m *Message) Command() string {
	if len(m.parts) > 0 {
		return m.parts[0]
	}
	return ""
}

// Params return the parsed message's params
func (m *Message) Params() []string {
	if len(m.parts) > 0 {
		return m.parts[1:]
	}

	return make([]string, 0)
}

func (m *Message) User() string {
	if name, ok := m.original.GetTag("display-name"); ok {
		return name
	}

	return m.original.Name
}

// Badges returns the badges from the tags if available
// It's in a <name>/<level> format
// Details: https://dev.twitch.tv/docs/irc/tags
func (m *Message) Badges() []string {
	if badge, ok := m.original.GetTag("badges"); ok {
		return strings.Split(badge, ",")
	}

	return make([]string, 0)
}
