package internal

import (
	"context"
	"fmt"

	twilio "github.com/twilio/twilio-go"
)

// twilioModule creates a Twilio REST client and registers it.
type twilioModule struct {
	name   string
	config map[string]any
}

func newTwilioModule(name string, config map[string]any) (*twilioModule, error) {
	return &twilioModule{name: name, config: config}, nil
}

// Init creates the Twilio REST client and registers it in the global registry.
func (m *twilioModule) Init() error {
	accountSid, _ := m.config["accountSid"].(string)
	authToken, _ := m.config["authToken"].(string)

	var client *twilio.RestClient

	if apiKey, ok := m.config["apiKey"].(string); ok && apiKey != "" {
		apiSecret, _ := m.config["apiSecret"].(string)
		params := twilio.ClientParams{
			Username:   apiKey,
			Password:   apiSecret,
			AccountSid: accountSid,
		}
		client = twilio.NewRestClientWithParams(params)
	} else {
		if accountSid == "" || authToken == "" {
			return fmt.Errorf("twilio.provider %q: accountSid and authToken are required", m.name)
		}
		params := twilio.ClientParams{
			Username:   accountSid,
			Password:   authToken,
			AccountSid: accountSid,
		}
		client = twilio.NewRestClientWithParams(params)
	}

	if region, ok := m.config["region"].(string); ok && region != "" {
		client.SetRegion(region)
	}
	if edge, ok := m.config["edge"].(string); ok && edge != "" {
		client.SetEdge(edge)
	}

	RegisterClient(m.name, client)
	return nil
}

// Start is a no-op for this module.
func (m *twilioModule) Start(_ context.Context) error { return nil }

// Stop unregisters the Twilio client.
func (m *twilioModule) Stop(_ context.Context) error {
	UnregisterClient(m.name)
	return nil
}
