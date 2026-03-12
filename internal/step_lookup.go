package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_lookup "github.com/twilio/twilio-go/rest/lookups/v2"
)

// lookupPhoneStep implements step.twilio_lookup_phone
type lookupPhoneStep struct {
	name       string
	moduleName string
}

func newLookupPhoneStep(name string, config map[string]any) (*lookupPhoneStep, error) {
	return &lookupPhoneStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *lookupPhoneStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	phoneNumber := resolveValue("phone_number", current, config)
	if phoneNumber == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "phone_number is required"}}, nil
	}
	params := &openapi_lookup.FetchPhoneNumberParams{}
	if fields := resolveValue("fields", current, config); fields != "" {
		params.SetFields(fields)
	}
	result, err := client.LookupsV2.FetchPhoneNumber(phoneNumber, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"phone_number":    derefStr(result.PhoneNumber),
		"country_code":    derefStr(result.CountryCode),
		"national_format": derefStr(result.NationalFormat),
		"valid":           derefBool(result.Valid),
	}}, nil
}
