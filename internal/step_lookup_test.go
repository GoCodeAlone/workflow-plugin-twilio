package internal

import (
	"context"
	"testing"
)

func TestLookupPhoneStep_MissingPhoneNumber(t *testing.T) {
	step, _ := newLookupPhoneStep("test", map[string]any{"module": "nonexistent-lookup"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing phone_number")
	}
}

func TestLookupPhoneStep_MissingClient(t *testing.T) {
	step, _ := newLookupPhoneStep("test", map[string]any{"module": "nonexistent-lookup2"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"phone_number": "+15555555555",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
