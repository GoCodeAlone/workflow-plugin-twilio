package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_verify "github.com/twilio/twilio-go/rest/verify/v2"
)

// sendVerificationStep implements step.twilio_send_verification
type sendVerificationStep struct {
	name       string
	moduleName string
}

func newSendVerificationStep(name string, config map[string]any) (*sendVerificationStep, error) {
	return &sendVerificationStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *sendVerificationStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	to := resolveValue("to", current, config)
	if to == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "to is required"}}, nil
	}
	params := &openapi_verify.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel(resolveValue("channel", current, config))
	v, err := client.VerifyV2.CreateVerification(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":     derefStr(v.Sid),
		"status":  derefStr(v.Status),
		"to":      derefStr(v.To),
		"channel": derefStr(v.Channel),
	}}, nil
}

// checkVerificationStep implements step.twilio_check_verification
type checkVerificationStep struct {
	name       string
	moduleName string
}

func newCheckVerificationStep(name string, config map[string]any) (*checkVerificationStep, error) {
	return &checkVerificationStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *checkVerificationStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	to := resolveValue("to", current, config)
	if to == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "to is required"}}, nil
	}
	code := resolveValue("code", current, config)
	if code == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "code is required"}}, nil
	}
	params := &openapi_verify.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)
	check, err := client.VerifyV2.CreateVerificationCheck(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"status": derefStr(check.Status),
		"valid":  derefBool(check.Valid),
		"to":     derefStr(check.To),
	}}, nil
}

// createVerifyServiceStep implements step.twilio_create_verify_service
type createVerifyServiceStep struct {
	name       string
	moduleName string
}

func newCreateVerifyServiceStep(name string, config map[string]any) (*createVerifyServiceStep, error) {
	return &createVerifyServiceStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createVerifyServiceStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi_verify.CreateServiceParams{}
	params.SetFriendlyName(friendlyName)
	svc, err := client.VerifyV2.CreateService(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(svc.Sid),
		"friendly_name": derefStr(svc.FriendlyName),
	}}, nil
}

// listVerifyServicesStep implements step.twilio_list_verify_services
type listVerifyServicesStep struct {
	name       string
	moduleName string
}

func newListVerifyServicesStep(name string, config map[string]any) (*listVerifyServicesStep, error) {
	return &listVerifyServicesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listVerifyServicesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	services, err := client.VerifyV2.ListService(&openapi_verify.ListServiceParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(services))
	for _, svc := range services {
		result = append(result, map[string]any{
			"sid":           derefStr(svc.Sid),
			"friendly_name": derefStr(svc.FriendlyName),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"services": result, "count": len(result)}}, nil
}
