package message_test

import (
	"testing"

	message "gh.iiji.jp/legs-v2/legs-message"

	"github.com/stretchr/testify/assert"
)

func TestNewConsoleStartMessage(t *testing.T) {
	msg := message.NewConsoleStart("session-id", "shell cmd")
	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Console{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "console", parsed.MessageType)
	assert.Equal(t, "console", parsed.Model)

	assert.Equal(t, "session-id", parsed.SessionID)
	assert.Equal(t, message.ConsoleStartState, parsed.State)
	assert.Equal(t, []byte("shell cmd"), parsed.Data)
}

func TestNewConsoleCloseMessage(t *testing.T) {
	msg := message.NewConsoleClose("session-id")
	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Console{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "console", parsed.MessageType)
	assert.Equal(t, "console", parsed.Model)

	assert.Equal(t, "session-id", parsed.SessionID)
	assert.Equal(t, message.ConsoleCloseState, parsed.State)
	assert.Equal(t, []byte(""), parsed.Data)
}

func TestNewConsoleInputMessage(t *testing.T) {
	msg := message.NewConsoleInput("session-id", []byte("test msg"))
	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Console{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "console", parsed.MessageType)
	assert.Equal(t, "console", parsed.Model)

	assert.Equal(t, "session-id", parsed.SessionID)
	assert.Equal(t, message.ConsoleInputState, parsed.State)
	assert.Equal(t, []byte("test msg"), parsed.Data)
}

func TestNewConsoleOutputMessage(t *testing.T) {
	msg := message.NewConsoleOutput("session-id", []byte("test msg"))
	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Console{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "console", parsed.MessageType)
	assert.Equal(t, "console", parsed.Model)

	assert.Equal(t, "session-id", parsed.SessionID)
	assert.Equal(t, message.ConsoleOutputState, parsed.State)
	assert.Equal(t, []byte("test msg"), parsed.Data)
}
