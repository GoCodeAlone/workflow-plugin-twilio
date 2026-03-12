package internal

import (
	"context"
	"testing"
)

func TestCreateAssistantStep_MissingName(t *testing.T) {
	step, _ := newCreateAssistantStep("test", map[string]any{"module": "nonexistent-asst"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing name")
	}
}

func TestListAssistantsStep_MissingClient(t *testing.T) {
	step, _ := newListAssistantsStep("test", map[string]any{"module": "nonexistent-listasst"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestCreateKnowledgeBaseStep_MissingClient(t *testing.T) {
	step, _ := newCreateKnowledgeBaseStep("test", map[string]any{"module": "nonexistent-kb"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
