package message

const (
	// TransferActionPost is the action for post json or text to external service.
	TransferActionPost = "post"

	// TransferActionGet is the action for download file from external service.
	TransferActionGet = "get"
)

// TransferRequestData defines request for http request transfer.
// Action: `TransferActionPost` or `TransferActionGet`
// Target: Transfer request destination. The Target can be a valid url (beginning with `http`) and routing name stored on the server.
// Value: The Request body for transfer http request.
type TransferRequestData struct {
	Action string `json:"action"`
	Target string `json:"target"`
	Value  string `json:"value"`
}

// TransferResponseData defines response when transferred http request.
// URL: transfer destination
// Body: http response body
// StatusCode: http response code
type TransferResponseData struct {
	URL        string `json:"url"`
	Body       []byte `json:"body"`
	StatusCode int    `json:"status_code"`
}

type transferData struct {
	SessionID string                 `json:"session_id"`
	Request   TransferRequestData    `json:"request"`
	Responses []TransferResponseData `json:"response"`
}

// Transfer defines message for data transfer request using raw url.
type Transfer struct {
	BaseMessage
	Data transferData `json:"data"`
}

// NewTransferRequest creates a transfer request message from string url.
func NewTransferRequest(sessionID string, req TransferRequestData) Message {
	data := transferData{
		SessionID: sessionID,
		Request:   req,
		Responses: []TransferResponseData{},
	}
	m := &Transfer{
		Data: data,
	}
	m.MessageType = "transfer"
	m.Model = "http_transfer"

	return m
}

// NewTransferResponse creates a transfer response message from request data.
func NewTransferResponse(data transferData, res []TransferResponseData) Message {
	data.Responses = res
	m := &Transfer{
		Data: data,
	}
	m.MessageType = "transfer"
	m.Model = "http_transfer"

	return m
}
