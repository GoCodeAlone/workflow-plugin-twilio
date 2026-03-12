package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_notify "github.com/twilio/twilio-go/rest/notify/v1"
)

// sendNotificationStep implements step.twilio_send_notification
type sendNotificationStep struct {
	name       string
	moduleName string
}

func newSendNotificationStep(name string, config map[string]any) (*sendNotificationStep, error) {
	return &sendNotificationStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *sendNotificationStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	params := &openapi_notify.CreateNotificationParams{}
	if body := resolveValue("body", current, config); body != "" {
		params.SetBody(body)
	}
	if title := resolveValue("title", current, config); title != "" {
		params.SetTitle(title)
	}
	if toBinding := resolveStringSlice("to_binding", current, config); len(toBinding) > 0 {
		params.SetToBinding(toBinding)
	}
	notif, err := client.NotifyV1.CreateNotification(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(notif.Sid),
		"account_sid": derefStr(notif.AccountSid),
	}}, nil
}

// createBindingStep implements step.twilio_create_binding
type createBindingStep struct {
	name       string
	moduleName string
}

func newCreateBindingStep(name string, config map[string]any) (*createBindingStep, error) {
	return &createBindingStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createBindingStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	identity := resolveValue("identity", current, config)
	if identity == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "identity is required"}}, nil
	}
	bindingType := resolveValue("binding_type", current, config)
	if bindingType == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "binding_type is required"}}, nil
	}
	address := resolveValue("address", current, config)
	if address == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "address is required"}}, nil
	}
	params := &openapi_notify.CreateBindingParams{}
	params.SetIdentity(identity)
	params.SetBindingType(bindingType)
	params.SetAddress(address)
	binding, err := client.NotifyV1.CreateBinding(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":          derefStr(binding.Sid),
		"identity":     derefStr(binding.Identity),
		"binding_type": derefStr(binding.BindingType),
	}}, nil
}

// listBindingsStep implements step.twilio_list_bindings
type listBindingsStep struct {
	name       string
	moduleName string
}

func newListBindingsStep(name string, config map[string]any) (*listBindingsStep, error) {
	return &listBindingsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listBindingsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	bindings, err := client.NotifyV1.ListBinding(serviceSid, &openapi_notify.ListBindingParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(bindings))
	for _, b := range bindings {
		result = append(result, map[string]any{
			"sid":          derefStr(b.Sid),
			"identity":     derefStr(b.Identity),
			"binding_type": derefStr(b.BindingType),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"bindings": result, "count": len(result)}}, nil
}

// createNotifyServiceStep implements step.twilio_create_notify_service
type createNotifyServiceStep struct {
	name       string
	moduleName string
}

func newCreateNotifyServiceStep(name string, config map[string]any) (*createNotifyServiceStep, error) {
	return &createNotifyServiceStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createNotifyServiceStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi_notify.CreateServiceParams{}
	params.SetFriendlyName(friendlyName)
	svc, err := client.NotifyV1.CreateService(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(svc.Sid),
		"friendly_name": derefStr(svc.FriendlyName),
	}}, nil
}
