package internal

import (
	"context"
	"testing"
)

func TestCreateServerlessServiceStep_MissingUniqueName(t *testing.T) {
	step, _ := newCreateServerlessServiceStep("test", map[string]any{"module": "nonexistent-sl"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing unique_name")
	}
}

func TestCreateFunctionStep_MissingServiceSid(t *testing.T) {
	step, _ := newCreateFunctionStep("test", map[string]any{"module": "nonexistent-fn"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}

func TestCreateBuildStep_MissingServiceSid(t *testing.T) {
	step, _ := newCreateBuildStep("test", map[string]any{"module": "nonexistent-build"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}
