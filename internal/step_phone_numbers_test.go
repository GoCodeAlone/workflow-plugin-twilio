package internal

import (
	"context"
	"testing"
)

func TestSearchAvailableStep_MissingCountryCode(t *testing.T) {
	step, _ := newSearchAvailableStep("test", map[string]any{"module": "nonexistent-search"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing country_code")
	}
}

func TestBuyNumberStep_MissingPhoneNumber(t *testing.T) {
	step, _ := newBuyNumberStep("test", map[string]any{"module": "nonexistent-buy"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing phone_number")
	}
}

func TestReleaseNumberStep_MissingSid(t *testing.T) {
	step, _ := newReleaseNumberStep("test", map[string]any{"module": "nonexistent-release"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing sid")
	}
}
