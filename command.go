package message

// CommandStatus defines type of command message status
type CommandStatus string

const (
	// CommandExecutingStatus specifies the status of executing command in client.
	CommandExecutingStatus CommandStatus = "executing"

	// CommandErrorStatus specifies the status of command execution was failed.
	CommandErrorStatus CommandStatus = "error"

	// CommandExecutedStatus specifies the status of command execution was success.
	CommandExecutedStatus CommandStatus = "executed"
)

// CommandMessage defines interface for command message methods
type CommandMessage interface {
	Message
	SetStatus(status CommandStatus)
	SetResult(result string)
}

// CommandData defines data for command message
type CommandData struct {
	ID       string        `json:"id"`
	DeviceID string        `json:"device_id"`
	Status   CommandStatus `json:"status"`
	Result   string        `json:"result"`
	Command  string        `json:"command"`
}

// Command defines message for command execution
type Command struct {
	BaseMessage
	Data CommandData `json:"data"`
}

// NewCommand creates command execution message
func NewCommand(commandID, deviceID string, status CommandStatus, result, command string) CommandMessage {
	data := CommandData{
		ID:       commandID,
		DeviceID: deviceID,
		Status:   status,
		Result:   result,
		Command:  command,
	}
	cm := &Command{
		Data: data,
	}
	cm.MessageType = "execute"
	cm.Model = "command"
	return cm
}

// SetStatus set a status attribute to command message.
func (c *Command) SetStatus(status CommandStatus) {
	c.Data.Status = status
}

// SetResult set a result of execution command to command message.
func (c *Command) SetResult(result string) {
	c.Data.Result = result
}
