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

func TestModuleInit_OptionalMissingCredentialsDoesNotRegisterClient(t *testing.T) {
	m, err := newTwilioModule("test-optional-missing", map[string]any{
		"optional": true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatalf("expected optional module with missing credentials to initialize, got %v", err)
	}
	if c, ok := GetClient("test-optional-missing"); ok || c != nil {
		t.Fatalf("expected no client to be registered for optional missing credentials, got ok=%v client=%v", ok, c)
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
