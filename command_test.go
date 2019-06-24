package message_test

import (
	"testing"

	message "gh.iiji.jp/legs-v2/legs-message"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandMessage(t *testing.T) {
	msg := message.NewCommand("cm-id", "device-id", message.CommandExecutingStatus, "result", "command")

	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Command{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "execute", parsed.MessageType)
	assert.Equal(t, "command", parsed.Model)

	assert.Equal(t, "cm-id", parsed.Data.ID)
	assert.Equal(t, "device-id", parsed.Data.DeviceID)
	assert.Equal(t, "executing", string(parsed.Data.Status))
	assert.Equal(t, "result", parsed.Data.Result)
	assert.Equal(t, "command", parsed.Data.Command)
}

func TestCommand_SetStatus(t *testing.T) {
	msg := message.NewCommand("cm-id", "device-id", message.CommandExecutingStatus, "result", "command")

	msg.SetStatus(message.CommandExecutedStatus)
	assert.Equal(t, "executed", string(msg.(*message.Command).Data.Status))
}

func TestCommand_SetResult(t *testing.T) {
	msg := message.NewCommand("cm-id", "device-id", message.CommandExecutingStatus, "result", "command")

	msg.SetResult("test result")
	assert.Equal(t, "test result", msg.(*message.Command).Data.Result)
}
