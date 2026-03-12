package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_sl "github.com/twilio/twilio-go/rest/serverless/v1"
)

// createServerlessServiceStep implements step.twilio_create_serverless_service
type createServerlessServiceStep struct {
	name       string
	moduleName string
}

func newCreateServerlessServiceStep(name string, config map[string]any) (*createServerlessServiceStep, error) {
	return &createServerlessServiceStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createServerlessServiceStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	uniqueName := resolveValue("unique_name", current, config)
	if uniqueName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "unique_name is required"}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		friendlyName = uniqueName
	}
	params := &openapi_sl.CreateServiceParams{}
	params.SetUniqueName(uniqueName)
	params.SetFriendlyName(friendlyName)
	svc, err := client.ServerlessV1.CreateService(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(svc.Sid),
		"unique_name":   derefStr(svc.UniqueName),
		"friendly_name": derefStr(svc.FriendlyName),
	}}, nil
}

// createFunctionStep implements step.twilio_create_function
type createFunctionStep struct {
	name       string
	moduleName string
}

func newCreateFunctionStep(name string, config map[string]any) (*createFunctionStep, error) {
	return &createFunctionStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createFunctionStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi_sl.CreateFunctionParams{}
	params.SetFriendlyName(friendlyName)
	fn, err := client.ServerlessV1.CreateFunction(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(fn.Sid),
		"friendly_name": derefStr(fn.FriendlyName),
	}}, nil
}

// createBuildStep implements step.twilio_create_build
type createBuildStep struct {
	name       string
	moduleName string
}

func newCreateBuildStep(name string, config map[string]any) (*createBuildStep, error) {
	return &createBuildStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createBuildStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	params := &openapi_sl.CreateBuildParams{}
	if fv := resolveStringSlice("function_versions", current, config); len(fv) > 0 {
		params.SetFunctionVersions(fv)
	}
	if av := resolveStringSlice("asset_versions", current, config); len(av) > 0 {
		params.SetAssetVersions(av)
	}
	if deps := resolveValue("dependencies", current, config); deps != "" {
		params.SetDependencies(deps)
	}
	build, err := client.ServerlessV1.CreateBuild(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(build.Sid),
		"status":      derefStr(build.Status),
		"service_sid": derefStr(build.ServiceSid),
	}}, nil
}

// listServerlessServicesStep implements step.twilio_list_serverless_services
type listServerlessServicesStep struct {
	name       string
	moduleName string
}

func newListServerlessServicesStep(name string, config map[string]any) (*listServerlessServicesStep, error) {
	return &listServerlessServicesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listServerlessServicesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	services, err := client.ServerlessV1.ListService(&openapi_sl.ListServiceParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(services))
	for _, svc := range services {
		result = append(result, map[string]any{
			"sid":           derefStr(svc.Sid),
			"unique_name":   derefStr(svc.UniqueName),
			"friendly_name": derefStr(svc.FriendlyName),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"services": result, "count": len(result)}}, nil
}
