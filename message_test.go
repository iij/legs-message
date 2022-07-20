package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	message "github.com/iij/legs-message"
)

func TestMarshal(t *testing.T) {
	msg := message.NewMessage()
	b, err := message.Marshal(msg)
	if err != nil {
		t.Error("error at ToBinary:", err)
		return
	}

	parsedMsg := message.NewMessage()
	err = message.Unmarshal(b, &parsedMsg)
	if err != nil {
		t.Error("error at ParseMessage:", err)
		return
	}

	if parsedMsg.MessageType != "none" {
		t.Error("message type was not 'none':", parsedMsg.MessageType)
	}
	if parsedMsg.Model != "none" {
		t.Error("message type was not 'none':", parsedMsg.Model)
	}
}

func TestGetMessageType(t *testing.T) {
	t.Skipf("RegisterExt is comment outed for backward compatibility")

	type test struct {
		input interface{}
		want  int8
		err   bool
	}

	cases := []test{
		{input: message.BaseMessage{}, want: message.TypeNone, err: false},
		{input: message.Command{}, want: message.TypeCommand, err: false},
		{input: message.ClientConfigure{}, want: message.TypeClientConfigure, err: false},
		{input: message.Console{}, want: message.TypeConsole, err: false},
		{input: message.Proxy{}, want: message.TypeProxy, err: false},
		{input: message.ProxyData{}, want: message.TypeInvalid, err: true},
	}

	for _, c := range cases {
		b, _ := message.Marshal(c.input)
		msgType, err := message.GetMessageType(b)
		if c.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, c.want, msgType)
	}
}
