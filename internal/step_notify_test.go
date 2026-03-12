package internal

import (
	"context"
	"testing"
)

func TestSendNotificationStep_MissingServiceSid(t *testing.T) {
	step, _ := newSendNotificationStep("test", map[string]any{"module": "nonexistent-notify"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}

func TestCreateBindingStep_MissingIdentity(t *testing.T) {
	step, _ := newCreateBindingStep("test", map[string]any{"module": "nonexistent-binding"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"service_sid": "IStest",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing identity")
	}
}

func TestCreateNotifyServiceStep_MissingFriendlyName(t *testing.T) {
	step, _ := newCreateNotifyServiceStep("test", map[string]any{"module": "nonexistent-notifysvc"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing friendly_name")
	}
}
