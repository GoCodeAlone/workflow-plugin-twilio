package internal

import (
	"context"
	"testing"
)

func TestSendSMSStep_MissingClient(t *testing.T) {
	step, err := newSendSMSStep("test", map[string]any{"module": "nonexistent-sms"})
	if err != nil {
		t.Fatal(err)
	}
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{"to": "+15555555555"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestSendSMSStep_MissingTo(t *testing.T) {
	step, _ := newSendSMSStep("test", map[string]any{"module": "nonexistent-sms2"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing to")
	}
}

func TestSendMMSStep_MissingClient(t *testing.T) {
	step, err := newSendMMSStep("test", map[string]any{"module": "nonexistent-mms"})
	if err != nil {
		t.Fatal(err)
	}
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{"to": "+15555555555"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestSendWhatsappStep_MissingTo(t *testing.T) {
	step, _ := newSendWhatsappStep("test", map[string]any{"module": "nonexistent-wa"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing to")
	}
}

func TestListMessagesStep_MissingClient(t *testing.T) {
	step, _ := newListMessagesStep("test", map[string]any{"module": "nonexistent-list"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestFetchMessageStep_MissingMessageSid(t *testing.T) {
	step, _ := newFetchMessageStep("test", map[string]any{"module": "nonexistent-fetch"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing message_sid")
	}
}

func TestDeleteMessageStep_MissingMessageSid(t *testing.T) {
	step, _ := newDeleteMessageStep("test", map[string]any{"module": "nonexistent-del"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing message_sid")
	}
}

func TestFetchMediaStep_MissingParams(t *testing.T) {
	step, _ := newFetchMediaStep("test", map[string]any{"module": "nonexistent-media"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing message_sid")
	}
}

func TestCreateMessagingServiceStep_MissingFriendlyName(t *testing.T) {
	step, _ := newCreateMessagingServiceStep("test", map[string]any{"module": "nonexistent-msgsvc"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing friendly_name")
	}
}
