package bot

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandRegister(t *testing.T) {
	b := New()

	b.Register("cmd1", nil)
	b.Register("cmd2", nil)
	b.Register("cmd2", nil)

	assert.Len(t, b.commands, 2)
}

func TestCommandHandlerFound(t *testing.T) {
	b := New()

	testError := errors.New("test error")
	var called bool

	b.Register("cmd1", nil)
	b.Register("cmd2", nil)
	b.Register("cmd2", func([]string) (string, error) {
		called = true
		return "resp1", testError
	})

	assert.Len(t, b.commands, 2)

	resp, err := b.Handler("cmd666", nil)
	assert.Equal(t, "", resp)
	assert.Equal(t, ErrNotFound, err)

	assert.False(t, called)
	resp, err = b.Handler("cmd2", nil)
	assert.Equal(t, "resp1", resp)
	assert.Equal(t, testError, err)
	assert.True(t, called)
}
