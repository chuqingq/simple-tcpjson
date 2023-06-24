# simple-tcpjson
A simple json client/server library on TCP.

## API

```go
type ClientState int

const (
	ClientConnected ClientState = iota
	ClientDisconnected
)

type OnClientStateChange func(c *Client, newstate ClientState)
type OnClientMsgRecv func(c *Client, msg *json.Json, err error)


// client
type Client {
}

func NewClient(serveraddr string) *Client
func (c *Client) SetTLS(key, cert, ca []byte) *Client
func (c *Client) SetOnStateChange(handler OnStateChange) *Client
func (c *Client) SetOnMsgRecv(handler OnMsgRecv) *Client
func (c *Client) Start() error
func (c *Client) Stop()

func (c *Client) Send(msg *json.Json) error
func (c *Client) State() ClientState


// server
type Server struct {
}

func NewServer(addr string) *Server
func (s *Server) SetTLS(key, cert, ca []byte) *Server
func (s *Server) SetOnPeerStateChange(handler OnClientStateChange) *Server
func (s *Server) SetOnMsgRecv(handler OnClientMsgRecv) *Server
func (s *Server) Start() error
func (s *Server) Stop()

```