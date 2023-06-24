package tcpjson

import json "github.com/chuqingq/simple-json"

type ClientState int

const (
	ClientConnected ClientState = iota
	ClientDisconnected
)

type OnClientStateChange func(c *Client, newstate ClientState)
type OnClientMsgRecv func(c *Client, msg *json.Json, err error)
