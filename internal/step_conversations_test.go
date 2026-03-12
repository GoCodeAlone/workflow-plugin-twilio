package internal

import (
	"context"
	"testing"
)

func TestCreateConversationStep_MissingClient(t *testing.T) {
	step, _ := newCreateConversationStep("test", map[string]any{"module": "nonexistent-conv"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestSendConversationMessageStep_MissingConversationSid(t *testing.T) {
	step, _ := newSendConversationMessageStep("test", map[string]any{"module": "nonexistent-convmsg"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing conversation_sid")
	}
}

func TestFetchConversationStep_MissingConversationSid(t *testing.T) {
	step, _ := newFetchConversationStep("test", map[string]any{"module": "nonexistent-fetchconv"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing conversation_sid")
	}
}

func TestCreateConversationUserStep_MissingIdentity(t *testing.T) {
	step, _ := newCreateConversationUserStep("test", map[string]any{"module": "nonexistent-convuser"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing identity")
	}
}
