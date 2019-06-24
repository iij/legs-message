package message

// ConsoleState defines type of console message state
type ConsoleState string

const (
	// ConsoleStartState is state of message of start console session to legs-client.
	ConsoleStartState ConsoleState = "start"

	// ConsoleInputState is state of message of input to legs-client console.
	ConsoleInputState ConsoleState = "input"

	// ConsoleCloseState is state of message of request for close console session.
	ConsoleCloseState ConsoleState = "close"

	// ConsoleOutputState is state of message of output from legs-client to server.
	ConsoleOutputState ConsoleState = "output"
)

// Console defines message for create console session between server and client.
type Console struct {
	BaseMessage
	State     ConsoleState `json:"state"`
	SessionID string       `json:"session_id"`
	Data      []byte       `json:"data"`
}

// NewConsoleStart creates console message at start session
func NewConsoleStart(sessionID, shell string) Message {
	return newConsole(sessionID, ConsoleStartState, []byte(shell))
}

// NewConsoleClose creates console message at close session.
func NewConsoleClose(sessionID string) Message {
	return newConsole(sessionID, ConsoleCloseState, []byte{})
}

// NewConsoleInput creates console message for input to client device.
func NewConsoleInput(sessionID string, in []byte) Message {
	return newConsole(sessionID, ConsoleInputState, in)
}

// NewConsoleOutput creates console message for output from client device.
func NewConsoleOutput(sessionID string, in []byte) Message {
	return newConsole(sessionID, ConsoleOutputState, in)
}

// NewConsole creates console message
func newConsole(sessionID string, state ConsoleState, data []byte) Message {
	cm := &Console{
		State:     state,
		Data:      data,
		SessionID: sessionID,
	}
	cm.MessageType = "console"
	cm.Model = "console"

	return cm
}
