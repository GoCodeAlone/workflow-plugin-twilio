package internal

import (
	"context"
	"testing"
)

func TestCreateCallStep_MissingTo(t *testing.T) {
	step, _ := newCreateCallStep("test", map[string]any{"module": "nonexistent-call"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing to")
	}
}

func TestFetchCallStep_MissingCallSid(t *testing.T) {
	step, _ := newFetchCallStep("test", map[string]any{"module": "nonexistent-fetch-call"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing call_sid")
	}
}

func TestListCallsStep_MissingClient(t *testing.T) {
	step, _ := newListCallsStep("test", map[string]any{"module": "nonexistent-list-calls"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestListConferencesStep_MissingClient(t *testing.T) {
	step, _ := newListConferencesStep("test", map[string]any{"module": "nonexistent-confs"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestAddParticipantStep_MissingConferenceSid(t *testing.T) {
	step, _ := newAddParticipantStep("test", map[string]any{"module": "nonexistent-participant"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing conference_sid")
	}
}

func TestCreateQueueStep_MissingFriendlyName(t *testing.T) {
	step, _ := newCreateQueueStep("test", map[string]any{"module": "nonexistent-queue"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing friendly_name")
	}
}

func TestFetchRecordingStep_MissingRecordingSid(t *testing.T) {
	step, _ := newFetchRecordingStep("test", map[string]any{"module": "nonexistent-rec"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing recording_sid")
	}
}
