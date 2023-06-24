package tcpjson

import (
	"log"
	"sync"
	"testing"

	sjon "github.com/chuqingq/simple-json"
	"github.com/stretchr/testify/assert"
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
	server.SetOnMsgRecv(func(peer *Client, msg *sjon.Json, err error) {
		log.Printf("peer[%p] recv: %v, err: %v", peer, msg, err)
		if err == nil {
			assert.Equal(t, msg.Get("name").MustString(), "value")
			wg.Done()
		}
	})
	err := server.Start()
	assert.Nil(t, err)
	defer server.Stop()

	client := NewClient(":12345")
	err = client.Start()
	assert.Nil(t, err)
	defer client.Stop()

	j := &sjon.Json{}
	j.Set("name", "value")

	err = client.Send(j)
	assert.Nil(t, err)
	wg.Wait()
}
