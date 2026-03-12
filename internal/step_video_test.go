package internal

import (
	"context"
	"testing"
)

func TestCreateRoomStep_MissingClient(t *testing.T) {
	step, _ := newCreateRoomStep("test", map[string]any{"module": "nonexistent-room"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestFetchRoomStep_MissingRoomSid(t *testing.T) {
	step, _ := newFetchRoomStep("test", map[string]any{"module": "nonexistent-fetchroom"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing room_sid")
	}
}

func TestCompleteRoomStep_MissingRoomSid(t *testing.T) {
	step, _ := newCompleteRoomStep("test", map[string]any{"module": "nonexistent-completeroom"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing room_sid")
	}
}

func TestCreateCompositionStep_MissingRoomSid(t *testing.T) {
	step, _ := newCreateCompositionStep("test", map[string]any{"module": "nonexistent-comp"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing room_sid")
	}
}
