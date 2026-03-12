package internal

import (
	"context"
	"testing"
)

func TestListSimsStep_MissingClient(t *testing.T) {
	step, _ := newListSimsStep("test", map[string]any{"module": "nonexistent-sims"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestFetchSimStep_MissingSimSid(t *testing.T) {
	step, _ := newFetchSimStep("test", map[string]any{"module": "nonexistent-fetchsim"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing sim_sid")
	}
}

func TestSendCommandStep_MissingCommand(t *testing.T) {
	step, _ := newSendCommandStep("test", map[string]any{"module": "nonexistent-cmd"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing command")
	}
}

func TestCreateRatePlanStep_MissingClient(t *testing.T) {
	step, _ := newCreateRatePlanStep("test", map[string]any{"module": "nonexistent-fleet"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
