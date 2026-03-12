package internal

import (
	"context"
	"testing"
)

func TestTriggerFlowStep_MissingFlowSid(t *testing.T) {
	step, _ := newTriggerFlowStep("test", map[string]any{"module": "nonexistent-studio"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing flow_sid")
	}
}

func TestListFlowsStep_MissingClient(t *testing.T) {
	step, _ := newListFlowsStep("test", map[string]any{"module": "nonexistent-flows"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestFetchExecutionStep_MissingExecutionSid(t *testing.T) {
	step, _ := newFetchExecutionStep("test", map[string]any{"module": "nonexistent-exec"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"flow_sid": "FWtest",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing execution_sid")
	}
}
