package message

// ClientConfigure defines message for configuring client basic config.
type ClientConfigure struct {
	BaseMessage
	Data ClientConfig `json:"data"`
}

// NewClientConfigure creates client configure message.
func NewClientConfigure(config ClientConfig) Message {
	m := &ClientConfigure{
		Data: config,
	}
	m.MessageType = "configure"
	m.Model = "client-config"

	return m
}

// ClientConfig defines data model for client basic setting.
type ClientConfig struct {
	PingInterval int    `json:"ping_interval"`
	DeviceID     string `json:"device_id"`
}
