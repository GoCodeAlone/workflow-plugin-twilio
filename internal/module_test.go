package internal

import (
	"context"
	"testing"
)

func TestModuleInit_RegistersClient(t *testing.T) {
	m, err := newTwilioModule("test-init", map[string]any{
		"accountSid": "ACtest",
		"authToken":  "tokentest",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatal(err)
	}
	c, ok := GetClient("test-init")
	if !ok || c == nil {
		t.Error("expected client to be registered")
	}
	// cleanup
	UnregisterClient("test-init")
}

func TestModuleStop_UnregistersClient(t *testing.T) {
	m, _ := newTwilioModule("test-stop", map[string]any{
		"accountSid": "ACtest",
		"authToken":  "tokentest",
	})
	_ = m.Init()
	_ = m.Stop(context.Background())
	_, ok := GetClient("test-stop")
	if ok {
		t.Error("expected client to be unregistered after stop")
	}
}

func TestModuleInit_MissingCredentials(t *testing.T) {
	m, err := newTwilioModule("test-missing", map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err == nil {
		t.Error("expected error for missing credentials")
		UnregisterClient("test-missing")
	}
}

func TestModuleInit_WithApiKey(t *testing.T) {
	m, err := newTwilioModule("test-apikey", map[string]any{
		"accountSid": "ACtest",
		"apiKey":     "SKtest",
		"apiSecret":  "secret",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatal(err)
	}
	c, ok := GetClient("test-apikey")
	if !ok || c == nil {
		t.Error("expected client to be registered with API key")
	}
	UnregisterClient("test-apikey")
}
