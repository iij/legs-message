package message

// ProxyData defines data for proxy message
type ProxyData struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Request   []byte `json:"request"`
	Response  []byte `json:"response"`
	ServerID  string `json:"server_id"`
	SessionID string `json:"session_id"`
}

// Proxy defines message for url forwarding protocol.
type Proxy struct {
	BaseMessage
	Data ProxyData `json:"data"`
}

// NewProxyRequest creates instance of Proxy with request bytes
func NewProxyRequest(serverID, id, sessionID, url string, req []byte) Message {
	data := ProxyData{
		ID:        id,
		Request:   req,
		URL:       url,
		ServerID:  serverID,
		SessionID: sessionID,
	}
	m := &Proxy{
		Data: data,
	}
	m.MessageType = "proxy"
	m.Model = "proxy_request"

	return m
}

// NewProxyResponse creates instance of Proxy with response bytes
func NewProxyResponse(requestData ProxyData, res []byte) Message {
	requestData.Response = res
	m := &Proxy{
		Data: requestData,
	}
	m.MessageType = "proxy"
	m.Model = "proxy_response"

	return m
}

// SetResponse set a response to proxy data
func (p *Proxy) SetResponse(response []byte) {
	p.Data.Response = response
}
