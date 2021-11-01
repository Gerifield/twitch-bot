package bot

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleIdeaSuccess(t *testing.T) {
	ts := &testSaver{}
	b := New(ts)

	assert.NoError(t, b.handleIdea("nagyon fasza otlet"))
	assert.NotEmpty(t, ts.paramID)
	assert.Equal(t, "nagyon fasza otlet", ts.paramText)
}

func TestHandleIdeaFail(t *testing.T) {
	testErr := errors.New("test error")
	ts := &testSaver{retError: testErr}
	b := New(ts)

	assert.Equal(t, testErr, b.handleIdea("nagyon fasza otlet"))
}

type testSaver struct {
	paramID   string
	paramText string

	retError error
}

func (ts *testSaver) Save(id string, text string) error {
	ts.paramID = id
	ts.paramText = text
	return ts.retError
}
