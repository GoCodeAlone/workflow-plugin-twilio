package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	openapi_pricing "github.com/twilio/twilio-go/rest/pricing/v2"
)

// fetchPricingStep implements step.twilio_fetch_pricing
type fetchPricingStep struct {
	name       string
	moduleName string
}

func newFetchPricingStep(name string, config map[string]any) (*fetchPricingStep, error) {
	return &fetchPricingStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchPricingStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	destinationNumber := resolveValue("destination_number", current, config)
	if destinationNumber == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "destination_number is required"}}, nil
	}
	params := &openapi_pricing.FetchVoiceNumberParams{}
	if originationNumber := resolveValue("origination_number", current, config); originationNumber != "" {
		params.SetOriginationNumber(originationNumber)
	}
	pricing, err := client.PricingV2.FetchVoiceNumber(destinationNumber, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	outboundPrices := make([]map[string]any, 0)
	if pricing.OutboundCallPrices != nil {
		for _, p := range *pricing.OutboundCallPrices {
			outboundPrices = append(outboundPrices, map[string]any{
				"base_price":    p.BasePrice,
				"current_price": p.CurrentPrice,
			})
		}
	}
	return &sdk.StepResult{Output: map[string]any{
		"destination_number":   derefStr(pricing.DestinationNumber),
		"origination_number":   derefStr(pricing.OriginationNumber),
		"outbound_call_prices": outboundPrices,
	}}, nil
}

// listUsageRecordsStep implements step.twilio_list_usage_records
type listUsageRecordsStep struct {
	name       string
	moduleName string
}

func newListUsageRecordsStep(name string, config map[string]any) (*listUsageRecordsStep, error) {
	return &listUsageRecordsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listUsageRecordsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi.ListUsageRecordParams{}
	if category := resolveValue("category", current, config); category != "" {
		params.SetCategory(category)
	}
	records, err := client.Api.ListUsageRecord(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(records))
	for _, r := range records {
		price := float32(0)
		if r.Price != nil {
			price = *r.Price
		}
		result = append(result, map[string]any{
			"category":    derefStr(r.Category),
			"description": derefStr(r.Description),
			"usage":       derefStr(r.Usage),
			"price":       price,
		})
	}
	return &sdk.StepResult{Output: map[string]any{"records": result, "count": len(result)}}, nil
}
