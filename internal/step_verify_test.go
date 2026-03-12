package internal

import (
	"context"
	"testing"
)

func TestSendVerificationStep_MissingServiceSid(t *testing.T) {
	step, _ := newSendVerificationStep("test", map[string]any{"module": "nonexistent-verify"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}

func TestCheckVerificationStep_MissingCode(t *testing.T) {
	step, _ := newCheckVerificationStep("test", map[string]any{"module": "nonexistent-check"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"service_sid": "VAtest",
		"to":          "+15555555555",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing code")
	}
}

func TestCreateVerifyServiceStep_MissingFriendlyName(t *testing.T) {
	step, _ := newCreateVerifyServiceStep("test", map[string]any{"module": "nonexistent-vsvc"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing friendly_name")
	}
}

func TestListVerifyServicesStep_MissingClient(t *testing.T) {
	step, _ := newListVerifyServicesStep("test", map[string]any{"module": "nonexistent-vlist"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
