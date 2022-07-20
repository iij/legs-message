package message_test

import (
	"testing"

	message "github.com/iij/legs-message"

	"github.com/stretchr/testify/assert"
)

func TestNewClientConfigureMessage(t *testing.T) {
	clientConfig := message.ClientConfig{
		PingInterval: 30,
		DeviceID:     "device-id",
	}

	msg := message.NewClientConfigure(clientConfig)
	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.ClientConfigure{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "configure", parsed.MessageType)
	assert.Equal(t, "client-config", parsed.Model)

	assert.Equal(t, clientConfig, parsed.Data)
}
