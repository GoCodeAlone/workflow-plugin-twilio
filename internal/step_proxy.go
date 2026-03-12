package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_proxy "github.com/twilio/twilio-go/rest/proxy/v1"
)

// createProxyServiceStep implements step.twilio_create_proxy_service
type createProxyServiceStep struct {
	name       string
	moduleName string
}

func newCreateProxyServiceStep(name string, config map[string]any) (*createProxyServiceStep, error) {
	return &createProxyServiceStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createProxyServiceStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	uniqueName := resolveValue("unique_name", current, config)
	if uniqueName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "unique_name is required"}}, nil
	}
	params := &openapi_proxy.CreateServiceParams{}
	params.SetUniqueName(uniqueName)
	svc, err := client.ProxyV1.CreateService(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(svc.Sid),
		"unique_name": derefStr(svc.UniqueName),
	}}, nil
}

// createSessionStep implements step.twilio_create_session
type createSessionStep struct {
	name       string
	moduleName string
}

func newCreateSessionStep(name string, config map[string]any) (*createSessionStep, error) {
	return &createSessionStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createSessionStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	params := &openapi_proxy.CreateSessionParams{}
	if uniqueName := resolveValue("unique_name", current, config); uniqueName != "" {
		params.SetUniqueName(uniqueName)
	}
	session, err := client.ProxyV1.CreateSession(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(session.Sid),
		"unique_name": derefStr(session.UniqueName),
		"status":      derefStr(session.Status),
	}}, nil
}

// addProxyParticipantStep implements step.twilio_add_proxy_participant
type addProxyParticipantStep struct {
	name       string
	moduleName string
}

func newAddProxyParticipantStep(name string, config map[string]any) (*addProxyParticipantStep, error) {
	return &addProxyParticipantStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *addProxyParticipantStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	sessionSid := resolveValue("session_sid", current, config)
	if sessionSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "session_sid is required"}}, nil
	}
	identifier := resolveValue("identifier", current, config)
	if identifier == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "identifier is required"}}, nil
	}
	params := &openapi_proxy.CreateParticipantParams{}
	params.SetIdentifier(identifier)
	if fn := resolveValue("friendly_name", current, config); fn != "" {
		params.SetFriendlyName(fn)
	}
	participant, err := client.ProxyV1.CreateParticipant(serviceSid, sessionSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":        derefStr(participant.Sid),
		"identifier": derefStr(participant.Identifier),
	}}, nil
}
