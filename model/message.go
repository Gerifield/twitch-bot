package model

import (
	"strings"

	"gopkg.in/irc.v3"
)

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
