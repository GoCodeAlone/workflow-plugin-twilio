package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// searchAvailableStep implements step.twilio_search_available
type searchAvailableStep struct {
	name       string
	moduleName string
}

func newSearchAvailableStep(name string, config map[string]any) (*searchAvailableStep, error) {
	return &searchAvailableStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *searchAvailableStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	countryCode := resolveValue("country_code", current, config)
	if countryCode == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "country_code is required"}}, nil
	}
	params := &openapi.ListAvailablePhoneNumberLocalParams{}
	if areaCode := resolveInt("area_code", current, config); areaCode != 0 {
		params.SetAreaCode(areaCode)
	}
	if contains := resolveValue("contains", current, config); contains != "" {
		params.SetContains(contains)
	}
	numbers, err := client.Api.ListAvailablePhoneNumberLocal(countryCode, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(numbers))
	for _, n := range numbers {
		result = append(result, map[string]any{
			"phone_number": derefStr(n.PhoneNumber),
			"friendly_name": derefStr(n.FriendlyName),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"numbers": result, "count": len(result)}}, nil
}

// buyNumberStep implements step.twilio_buy_number
type buyNumberStep struct {
	name       string
	moduleName string
}

func newBuyNumberStep(name string, config map[string]any) (*buyNumberStep, error) {
	return &buyNumberStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *buyNumberStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	phoneNumber := resolveValue("phone_number", current, config)
	if phoneNumber == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "phone_number is required"}}, nil
	}
	params := &openapi.CreateIncomingPhoneNumberParams{}
	params.SetPhoneNumber(phoneNumber)
	if fn := resolveValue("friendly_name", current, config); fn != "" {
		params.SetFriendlyName(fn)
	}
	num, err := client.Api.CreateIncomingPhoneNumber(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(num.Sid),
		"phone_number":  derefStr(num.PhoneNumber),
		"friendly_name": derefStr(num.FriendlyName),
	}}, nil
}

// listNumbersStep implements step.twilio_list_numbers
type listNumbersStep struct {
	name       string
	moduleName string
}

func newListNumbersStep(name string, config map[string]any) (*listNumbersStep, error) {
	return &listNumbersStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listNumbersStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	numbers, err := client.Api.ListIncomingPhoneNumber(&openapi.ListIncomingPhoneNumberParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(numbers))
	for _, n := range numbers {
		result = append(result, map[string]any{
			"sid":           derefStr(n.Sid),
			"phone_number":  derefStr(n.PhoneNumber),
			"friendly_name": derefStr(n.FriendlyName),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"numbers": result, "count": len(result)}}, nil
}

// updateNumberStep implements step.twilio_update_number
type updateNumberStep struct {
	name       string
	moduleName string
}

func newUpdateNumberStep(name string, config map[string]any) (*updateNumberStep, error) {
	return &updateNumberStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *updateNumberStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "sid is required"}}, nil
	}
	params := &openapi.UpdateIncomingPhoneNumberParams{}
	if fn := resolveValue("friendly_name", current, config); fn != "" {
		params.SetFriendlyName(fn)
	}
	if smsUrl := resolveValue("sms_url", current, config); smsUrl != "" {
		params.SetSmsUrl(smsUrl)
	}
	if voiceUrl := resolveValue("voice_url", current, config); voiceUrl != "" {
		params.SetVoiceUrl(voiceUrl)
	}
	num, err := client.Api.UpdateIncomingPhoneNumber(sid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(num.Sid),
		"phone_number":  derefStr(num.PhoneNumber),
		"friendly_name": derefStr(num.FriendlyName),
	}}, nil
}

// releaseNumberStep implements step.twilio_release_number
type releaseNumberStep struct {
	name       string
	moduleName string
}

func newReleaseNumberStep(name string, config map[string]any) (*releaseNumberStep, error) {
	return &releaseNumberStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *releaseNumberStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "sid is required"}}, nil
	}
	err := client.Api.DeleteIncomingPhoneNumber(sid, &openapi.DeleteIncomingPhoneNumberParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true}}, nil
}
