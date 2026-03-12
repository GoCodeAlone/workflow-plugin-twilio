package internal

import (
	"context"
	"testing"
)

func TestCreateTrustProductStep_MissingEmail(t *testing.T) {
	step, _ := newCreateTrustProductStep("test", map[string]any{"module": "nonexistent-trust"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"friendly_name": "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing email")
	}
}

func TestListTrustProductsStep_MissingClient(t *testing.T) {
	step, _ := newListTrustProductsStep("test", map[string]any{"module": "nonexistent-listtrust"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestFetchTrustProductStep_MissingSid(t *testing.T) {
	step, _ := newFetchTrustProductStep("test", map[string]any{"module": "nonexistent-fetchtrust"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing sid")
	}
}
