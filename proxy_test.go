package message_test

import (
	"testing"

	message "gh.iiji.jp/legs-v2/legs-message"

	"github.com/stretchr/testify/assert"
)

func TestNewProxyRequest(t *testing.T) {
	msg := message.NewProxyRequest("server-id", "id", "session-id", "url", []byte("request"))
	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Proxy{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "proxy", parsed.MessageType)
	assert.Equal(t, "proxy_request", parsed.Model)

	assert.Equal(t, "server-id", parsed.Data.ServerID)
	assert.Equal(t, "id", parsed.Data.ID)
	assert.Equal(t, "session-id", parsed.Data.SessionID)
	assert.Equal(t, "url", parsed.Data.URL)
	assert.Equal(t, []byte("request"), parsed.Data.Request)
	assert.Equal(t, []byte(nil), parsed.Data.Response)
}

func TestNewProxyResponse(t *testing.T) {
	requestMsg := message.NewProxyRequest("server-id", "id", "session-id", "url", []byte("request"))

	msg := message.NewProxyResponse(requestMsg.(*message.Proxy).Data, []byte("response message"))
	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Proxy{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "proxy", parsed.MessageType)
	assert.Equal(t, "proxy_response", parsed.Model)

	assert.Equal(t, requestMsg.(*message.Proxy).Data.Request, parsed.Data.Request)
	assert.Equal(t, []byte("response message"), parsed.Data.Response)
}
