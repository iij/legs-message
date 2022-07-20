package message_test

import (
	"testing"

	message "github.com/iij/legs-message"
	"github.com/stretchr/testify/assert"
)

func TestNewTransferRequestURL(t *testing.T) {
	sessionID := "session-id"
	value := "transfer value"
	action := message.TransferActionPost
	url := "test-url"

	req := message.TransferRequestData{
		Target: url,
		Value:  value,
		Action: action,
	}
	msg := message.NewTransferRequest(sessionID, req)
	assert.Equal(t, "transfer", msg.GetMessageType())
	assert.Equal(t, "http_transfer", msg.GetModel())

	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Transfer{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "transfer", parsed.MessageType)
	assert.Equal(t, "http_transfer", parsed.Model)

	assert.Equal(t, action, parsed.Data.Request.Action)
	assert.Equal(t, sessionID, parsed.Data.SessionID)
	assert.Equal(t, value, parsed.Data.Request.Value)
	assert.Equal(t, []message.TransferResponseData{}, parsed.Data.Responses)
	assert.Equal(t, url, parsed.Data.Request.Target)
}

func TestNewTransferResponse(t *testing.T) {
	sessionID := "session-id"
	value := "transfer value"
	action := message.TransferActionGet
	url := "test-url"
	req := message.TransferRequestData{
		Target: url,
		Value:  value,
		Action: action,
	}
	responses := []message.TransferResponseData{
		{URL: "url", Body: []byte("response body"), StatusCode: 200},
		{URL: "url2", Body: []byte("response body"), StatusCode: 200},
	}
	reqMsg := message.NewTransferRequest(sessionID, req)

	msg := message.NewTransferResponse(reqMsg.(*message.Transfer).Data, responses)
	assert.Equal(t, "transfer", msg.GetMessageType())
	assert.Equal(t, "http_transfer", msg.GetModel())

	b, err := message.Marshal(msg)
	assert.Nil(t, err)

	parsed := &message.Transfer{}
	err = message.Unmarshal(b, parsed)
	assert.Nil(t, err)

	assert.Equal(t, "transfer", parsed.MessageType)
	assert.Equal(t, "http_transfer", parsed.Model)

	assert.Equal(t, action, parsed.Data.Request.Action)
	assert.Equal(t, sessionID, parsed.Data.SessionID)
	assert.Equal(t, value, parsed.Data.Request.Value)
	assert.Equal(t, responses, parsed.Data.Responses)
	assert.Equal(t, url, parsed.Data.Request.Target)
}
