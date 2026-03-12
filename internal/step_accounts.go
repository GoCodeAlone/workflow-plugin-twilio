package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// listAccountsStep implements step.twilio_list_accounts
type listAccountsStep struct {
	name       string
	moduleName string
}

func newListAccountsStep(name string, config map[string]any) (*listAccountsStep, error) {
	return &listAccountsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listAccountsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	accounts, err := client.Api.ListAccount(&openapi.ListAccountParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(accounts))
	for _, a := range accounts {
		result = append(result, map[string]any{
			"sid":           derefStr(a.Sid),
			"friendly_name": derefStr(a.FriendlyName),
			"status":        derefStr(a.Status),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"accounts": result, "count": len(result)}}, nil
}

// createApiKeyStep implements step.twilio_create_api_key
type createApiKeyStep struct {
	name       string
	moduleName string
}

func newCreateApiKeyStep(name string, config map[string]any) (*createApiKeyStep, error) {
	return &createApiKeyStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createApiKeyStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi.CreateNewKeyParams{}
	if fn := resolveValue("friendly_name", current, config); fn != "" {
		params.SetFriendlyName(fn)
	}
	key, err := client.Api.CreateNewKey(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(key.Sid),
		"friendly_name": derefStr(key.FriendlyName),
		"secret":        derefStr(key.Secret),
	}}, nil
}

// listApiKeysStep implements step.twilio_list_api_keys
type listApiKeysStep struct {
	name       string
	moduleName string
}

func newListApiKeysStep(name string, config map[string]any) (*listApiKeysStep, error) {
	return &listApiKeysStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listApiKeysStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	keys, err := client.Api.ListKey(&openapi.ListKeyParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(keys))
	for _, k := range keys {
		result = append(result, map[string]any{
			"sid":           derefStr(k.Sid),
			"friendly_name": derefStr(k.FriendlyName),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"keys": result, "count": len(result)}}, nil
}
