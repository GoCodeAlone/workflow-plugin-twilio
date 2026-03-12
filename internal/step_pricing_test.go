package internal

import (
	"context"
	"testing"
)

func TestFetchPricingStep_MissingDestinationNumber(t *testing.T) {
	step, _ := newFetchPricingStep("test", map[string]any{"module": "nonexistent-pricing"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing destination_number")
	}
}

func TestListUsageRecordsStep_MissingClient(t *testing.T) {
	step, _ := newListUsageRecordsStep("test", map[string]any{"module": "nonexistent-usage"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}
