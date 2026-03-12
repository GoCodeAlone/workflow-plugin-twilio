package internal

import (
	"context"
	"testing"
)

func TestCreateFlexFlowStep_MissingChannelType(t *testing.T) {
	step, _ := newCreateFlexFlowStep("test", map[string]any{"module": "nonexistent-flex"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"friendly_name": "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing channel_type")
	}
}

func TestCreateWebChannelStep_MissingFlexFlowSid(t *testing.T) {
	step, _ := newCreateWebChannelStep("test", map[string]any{"module": "nonexistent-webchan"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing flex_flow_sid")
	}
}

func TestListFlexFlowsStep_MissingClient(t *testing.T) {
	step, _ := newListFlexFlowsStep("test", map[string]any{"module": "nonexistent-listflex"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
