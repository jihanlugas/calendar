package ws

import (
	"sync"

	"github.com/labstack/echo/v4"
)

type HubManager struct {
	hubs      map[string]*PropertyHub
	mu        sync.RWMutex
	Validator echo.Validator
}

func NewHubManager(validator echo.Validator) *HubManager {
	return &HubManager{
		hubs:      make(map[string]*PropertyHub),
		Validator: validator,
	}
}

func (m *HubManager) GetHub(propertyID string) *PropertyHub {
	m.mu.RLock()
	hub, ok := m.hubs[propertyID]
	m.mu.RUnlock()

	if ok {
		return hub
	}

	m.mu.Lock()
	hub = NewPropertyHub(propertyID)
	m.hubs[propertyID] = hub
	m.mu.Unlock()

	go hub.Run()
	return hub
}

func (m *HubManager) SendToProperty(propertyID string, msg []byte) {
	hub := m.GetHub(propertyID)
	hub.broadcast <- msg
}
