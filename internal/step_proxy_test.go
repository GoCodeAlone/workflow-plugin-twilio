package internal

import (
	"context"
	"testing"
)

func TestCreateProxyServiceStep_MissingUniqueName(t *testing.T) {
	step, _ := newCreateProxyServiceStep("test", map[string]any{"module": "nonexistent-proxy"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing unique_name")
	}
}

func TestCreateSessionStep_MissingServiceSid(t *testing.T) {
	step, _ := newCreateSessionStep("test", map[string]any{"module": "nonexistent-session"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}

func TestAddProxyParticipantStep_MissingIdentifier(t *testing.T) {
	step, _ := newAddProxyParticipantStep("test", map[string]any{"module": "nonexistent-proxyp"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"service_sid": "KStest",
		"session_sid": "KCtest",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing identifier")
	}
}
