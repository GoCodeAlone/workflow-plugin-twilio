package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_studio "github.com/twilio/twilio-go/rest/studio/v2"
)

// triggerFlowStep implements step.twilio_trigger_flow
type triggerFlowStep struct {
	name       string
	moduleName string
}

func newTriggerFlowStep(name string, config map[string]any) (*triggerFlowStep, error) {
	return &triggerFlowStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *triggerFlowStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	flowSid := resolveValue("flow_sid", current, config)
	if flowSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "flow_sid is required"}}, nil
	}
	params := &openapi_studio.CreateExecutionParams{}
	params.SetTo(resolveValue("to", current, config))
	params.SetFrom(resolveValue("from", current, config))
	exec, err := client.StudioV2.CreateExecution(flowSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":      derefStr(exec.Sid),
		"status":   derefStr(exec.Status),
		"flow_sid": derefStr(exec.FlowSid),
	}}, nil
}

// listFlowsStep implements step.twilio_list_flows
type listFlowsStep struct {
	name       string
	moduleName string
}

func newListFlowsStep(name string, config map[string]any) (*listFlowsStep, error) {
	return &listFlowsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listFlowsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	flows, err := client.StudioV2.ListFlow(&openapi_studio.ListFlowParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(flows))
	for _, f := range flows {
		result = append(result, map[string]any{
			"sid":           derefStr(f.Sid),
			"friendly_name": derefStr(f.FriendlyName),
			"status":        derefStr(f.Status),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"flows": result, "count": len(result)}}, nil
}

// fetchExecutionStep implements step.twilio_fetch_execution
type fetchExecutionStep struct {
	name       string
	moduleName string
}

func newFetchExecutionStep(name string, config map[string]any) (*fetchExecutionStep, error) {
	return &fetchExecutionStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchExecutionStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	flowSid := resolveValue("flow_sid", current, config)
	if flowSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "flow_sid is required"}}, nil
	}
	executionSid := resolveValue("execution_sid", current, config)
	if executionSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "execution_sid is required"}}, nil
	}
	exec, err := client.StudioV2.FetchExecution(flowSid, executionSid)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(exec.Sid),
		"status":      derefStr(exec.Status),
		"flow_sid":    derefStr(exec.FlowSid),
		"account_sid": derefStr(exec.AccountSid),
	}}, nil
}
