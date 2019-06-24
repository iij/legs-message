package message

import (
	"bytes"

	"github.com/vmihailenco/msgpack"
)

// Message defines interface for every message structs
type Message interface {
	GetMessageType() string
	GetModel() string
}

// NewMessage creates instance of plane message
func NewMessage() BaseMessage {
	return BaseMessage{
		MessageType: "none",
		Model:       "none",
	}
}

// BaseMessage defines basic message.
type BaseMessage struct {
	MessageType string `json:"messageType"`
	Model       string `json:"model"`
}

// GetMessageType returns message type property.
func (m *BaseMessage) GetMessageType() string {
	return m.MessageType
}

// GetModel returns model property.
func (m *BaseMessage) GetModel() string {
	return m.Model
}

// Unmarshal returns decoding msgpack with type of msg param.
func Unmarshal(b []byte, msg interface{}) error {
	decoder := msgpack.NewDecoder(bytes.NewReader(b))
	decoder.UseJSONTag(true)

	return decoder.Decode(msg)
}

// Marshal returns msgpack encoding of msg
func Marshal(msg interface{}) (packBytes []byte, err error) {
	var buff bytes.Buffer
	encoder := msgpack.NewEncoder(&buff)
	encoder.UseJSONTag(true)

	err = encoder.Encode(msg)
	packBytes = buff.Bytes()

	return
}

// GetMessageType gets type for message by extended header in msgpack
func GetMessageType(msg []byte) (id int8, err error) {
	decoder := msgpack.NewDecoder(bytes.NewReader(msg))
	id, _, err = decoder.DecodeExtHeader()

	return
}
