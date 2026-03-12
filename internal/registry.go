package internal

import (
	"sync"

	twilio "github.com/twilio/twilio-go"
)

var (
	clientMu       sync.RWMutex
	clientRegistry = make(map[string]*twilio.RestClient)
)

// RegisterClient adds a Twilio REST client to the global registry under the given name.
func RegisterClient(name string, c *twilio.RestClient) {
	clientMu.Lock()
	defer clientMu.Unlock()
	clientRegistry[name] = c
}

// GetClient looks up a Twilio REST client by name.
func GetClient(name string) (*twilio.RestClient, bool) {
	clientMu.RLock()
	defer clientMu.RUnlock()
	c, ok := clientRegistry[name]
	return c, ok
}

// UnregisterClient removes a client from the registry.
func UnregisterClient(name string) {
	clientMu.Lock()
	defer clientMu.Unlock()
	delete(clientRegistry, name)
}
