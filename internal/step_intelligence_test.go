package internal

import (
	"context"
	"testing"
)

func TestCreateTranscriptStep_MissingServiceSid(t *testing.T) {
	step, _ := newCreateTranscriptStep("test", map[string]any{"module": "nonexistent-intel"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}

func TestFetchTranscriptStep_MissingSid(t *testing.T) {
	step, _ := newFetchTranscriptStep("test", map[string]any{"module": "nonexistent-fetcht"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing sid")
	}
}

func TestListTranscriptsStep_MissingClient(t *testing.T) {
	step, _ := newListTranscriptsStep("test", map[string]any{"module": "nonexistent-listt"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
