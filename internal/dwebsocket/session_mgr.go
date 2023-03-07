package dwebsocket

import (
	"db-go-gin/internal/global"
	"encoding/json"
	"go.uber.org/zap"
	"sync"
)

var clientManager = NewClientManager()

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients: make(map[*Client]struct{}),
	}
}

type ClientManager struct {
	Clients map[*Client]struct{}
	Lock    sync.RWMutex
}

func Add(client *Client) {
	clientManager.Lock.Lock()
	defer clientManager.Lock.Unlock()

	clientManager.Clients[client] = struct{}{}
}

func Delete(client *Client) {
	clientManager.Lock.Lock()
	defer clientManager.Lock.Unlock()

	delete(clientManager.Clients, client)
}

func GetAll() []*Client {
	clientManager.Lock.RLock()
	defer clientManager.Lock.RUnlock()

	var client []*Client

	for key, _ := range clientManager.Clients {
		client = append(client, key)
	}

	return client
}

// SendAll 通知所有前端
func SendAll(data interface{}) {
	msg, err := json.Marshal(data)
	if err != nil {
		global.LOG.Error("ws json marshal error", zap.Error(err))
		return
	}

	for _, value := range GetAll() {
		value.Out <- msg
	}
}
