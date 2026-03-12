package internal

import (
	"context"
	"testing"
)

func TestCreateSyncServiceStep_MissingClient(t *testing.T) {
	step, _ := newCreateSyncServiceStep("test", map[string]any{"module": "nonexistent-sync"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestCreateDocumentStep_MissingServiceSid(t *testing.T) {
	step, _ := newCreateDocumentStep("test", map[string]any{"module": "nonexistent-doc"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}

func TestUpdateDocumentStep_MissingDocumentSid(t *testing.T) {
	step, _ := newUpdateDocumentStep("test", map[string]any{"module": "nonexistent-updatedoc"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"service_sid": "IStest",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing document_sid")
	}
}

func TestCreateSyncMapStep_MissingServiceSid(t *testing.T) {
	step, _ := newCreateSyncMapStep("test", map[string]any{"module": "nonexistent-syncmap"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing service_sid")
	}
}
