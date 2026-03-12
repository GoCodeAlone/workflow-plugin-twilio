package internal

import (
	"context"
	"testing"
)

func TestListAccountsStep_MissingClient(t *testing.T) {
	step, _ := newListAccountsStep("test", map[string]any{"module": "nonexistent-accts"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestCreateApiKeyStep_MissingClient(t *testing.T) {
	step, _ := newCreateApiKeyStep("test", map[string]any{"module": "nonexistent-apikey"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestListApiKeysStep_MissingClient(t *testing.T) {
	step, _ := newListApiKeysStep("test", map[string]any{"module": "nonexistent-listkeys"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
