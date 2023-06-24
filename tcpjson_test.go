package tcpjson

import (
	"log"
	"sync"
	"testing"

	"github.com/bmizerany/assert"
	json "github.com/chuqingq/simple-json"
)

func TestMain(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	server := NewServer(":12345")
	server.SetOnPeerStateChange(func(peer *Client, state ClientState) {
		log.Printf("peer[%p] state change: %v", peer, state)
		if state == ClientConnected {
			wg.Done()
		}
	})
	server.SetOnMsgRecv(func(peer *Client, msg *json.Json, err error) {
		log.Printf("peer[%p] recv: %v, err: %v", peer, msg, err)
		if err == nil {
			assert.Equal(t, msg.Get("name").MustString(), "value")
			wg.Done()
		}
	})
	server.Start()
	defer server.Stop()

	client := NewClient(":12345")
	client.Start()
	defer client.Stop()

	j := &json.Json{}
	j.Set("name", "value")

	client.Send(j)
	wg.Wait()
}
