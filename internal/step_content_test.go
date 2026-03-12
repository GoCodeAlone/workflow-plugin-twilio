package internal

import (
	"context"
	"testing"
)

func TestCreateContentTemplateStep_MissingFriendlyName(t *testing.T) {
	step, _ := newCreateContentTemplateStep("test", map[string]any{"module": "nonexistent-content"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing friendly_name")
	}
}

func TestCreateContentTemplateStep_MissingLanguage(t *testing.T) {
	step, _ := newCreateContentTemplateStep("test", map[string]any{"module": "nonexistent-content2"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{
		"friendly_name": "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing language")
	}
}

func TestFetchContentTemplateStep_MissingSid(t *testing.T) {
	step, _ := newFetchContentTemplateStep("test", map[string]any{"module": "nonexistent-fetchcontent"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing sid")
	}
}
