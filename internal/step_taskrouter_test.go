package internal

import (
	"context"
	"testing"
)

func TestCreateWorkspaceStep_MissingFriendlyName(t *testing.T) {
	step, _ := newCreateWorkspaceStep("test", map[string]any{"module": "nonexistent-ws"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing friendly_name")
	}
}

func TestCreateTaskStep_MissingWorkspaceSid(t *testing.T) {
	step, _ := newCreateTaskStep("test", map[string]any{"module": "nonexistent-task"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing workspace_sid")
	}
}

func TestUpdateTaskStep_MissingTaskSid(t *testing.T) {
	step, _ := newUpdateTaskStep("test", map[string]any{"module": "nonexistent-updatetask"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"workspace_sid": "WStest",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing task_sid")
	}
}

func TestCreateTRWorkflowStep_MissingConfiguration(t *testing.T) {
	step, _ := newCreateTRWorkflowStep("test", map[string]any{"module": "nonexistent-trwf"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"workspace_sid": "WStest",
		"friendly_name": "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing configuration")
	}
}
