# simple-tcpjson
A simple json client/server library on TCP.

## go doc

```
package tcpjson // import "tcpjson"


TYPES

type Client struct {
	ServerAddr string
	Key        []byte
	Cert       []byte
	Ca         []byte

	OnStateChange OnClientStateChange
	OnMsgRecv     OnClientMsgRecv

	Conn net.Conn

	State ClientState
	// Has unexported fields.
}

func NewClient(serveraddr string) *Client

func (c *Client) Send(msg *sjson.Json) error

func (c *Client) SetOnMsgRecv(handler OnClientMsgRecv) *Client

func (c *Client) SetOnStateChange(handler OnClientStateChange) *Client

func (c *Client) SetTLS(key, cert, ca []byte) *Client

func (c *Client) Start() error

func (c *Client) Stop()

type ClientState int

const (
	ClientConnected ClientState = iota
	ClientDisconnected
)
type OnClientMsgRecv func(c *Client, msg *json.Json, err error)

type OnClientStateChange func(c *Client, newstate ClientState)

type Server struct {
	ServerAddr string
	Key        []byte
	Cert       []byte
	Ca         []byte

	OnClientStateChange OnClientStateChange
	OnClientMsgRecv     OnClientMsgRecv

	Listener net.Listener
	Conns    []net.Conn
}

func NewServer(addr string) *Server

func (s *Server) SetOnMsgRecv(handler OnClientMsgRecv) *Server

func (s *Server) SetOnPeerStateChange(handler OnClientStateChange) *Server

func (s *Server) SetTLS(key, cert, ca []byte) *Server

func (s *Server) Start() error

func (s *Server) Stop()
```
