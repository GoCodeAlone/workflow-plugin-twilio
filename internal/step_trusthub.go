package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_th "github.com/twilio/twilio-go/rest/trusthub/v1"
)

// createTrustProductStep implements step.twilio_create_trust_product
type createTrustProductStep struct {
	name       string
	moduleName string
}

func newCreateTrustProductStep(name string, config map[string]any) (*createTrustProductStep, error) {
	return &createTrustProductStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createTrustProductStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	email := resolveValue("email", current, config)
	if email == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "email is required"}}, nil
	}
	policySid := resolveValue("policy_sid", current, config)
	if policySid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "policy_sid is required"}}, nil
	}
	params := &openapi_th.CreateTrustProductParams{}
	params.SetFriendlyName(friendlyName)
	params.SetEmail(email)
	params.SetPolicySid(policySid)
	product, err := client.TrusthubV1.CreateTrustProduct(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(product.Sid),
		"friendly_name": derefStr(product.FriendlyName),
		"status":        derefStr(product.Status),
	}}, nil
}

// listTrustProductsStep implements step.twilio_list_trust_products
type listTrustProductsStep struct {
	name       string
	moduleName string
}

func newListTrustProductsStep(name string, config map[string]any) (*listTrustProductsStep, error) {
	return &listTrustProductsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listTrustProductsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	products, err := client.TrusthubV1.ListTrustProduct(&openapi_th.ListTrustProductParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(products))
	for _, p := range products {
		result = append(result, map[string]any{
			"sid":           derefStr(p.Sid),
			"friendly_name": derefStr(p.FriendlyName),
			"status":        derefStr(p.Status),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"products": result, "count": len(result)}}, nil
}

// fetchTrustProductStep implements step.twilio_fetch_trust_product
type fetchTrustProductStep struct {
	name       string
	moduleName string
}

func newFetchTrustProductStep(name string, config map[string]any) (*fetchTrustProductStep, error) {
	return &fetchTrustProductStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchTrustProductStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "sid is required"}}, nil
	}
	product, err := client.TrusthubV1.FetchTrustProduct(sid)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(product.Sid),
		"friendly_name": derefStr(product.FriendlyName),
		"status":        derefStr(product.Status),
	}}, nil
}
