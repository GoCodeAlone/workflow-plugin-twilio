package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_wireless "github.com/twilio/twilio-go/rest/wireless/v1"
)

// listSimsStep implements step.twilio_list_sims
type listSimsStep struct {
	name       string
	moduleName string
}

func newListSimsStep(name string, config map[string]any) (*listSimsStep, error) {
	return &listSimsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listSimsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sims, err := client.WirelessV1.ListSim(&openapi_wireless.ListSimParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(sims))
	for _, sim := range sims {
		result = append(result, map[string]any{
			"sid":         derefStr(sim.Sid),
			"unique_name": derefStr(sim.UniqueName),
			"status":      derefStr(sim.Status),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"sims": result, "count": len(result)}}, nil
}

// fetchSimStep implements step.twilio_fetch_sim
type fetchSimStep struct {
	name       string
	moduleName string
}

func newFetchSimStep(name string, config map[string]any) (*fetchSimStep, error) {
	return &fetchSimStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchSimStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("sim_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "sim_sid is required"}}, nil
	}
	sim, err := client.WirelessV1.FetchSim(sid)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(sim.Sid),
		"unique_name": derefStr(sim.UniqueName),
		"status":      derefStr(sim.Status),
	}}, nil
}

// updateSimStep implements step.twilio_update_sim
type updateSimStep struct {
	name       string
	moduleName string
}

func newUpdateSimStep(name string, config map[string]any) (*updateSimStep, error) {
	return &updateSimStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *updateSimStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("sim_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "sim_sid is required"}}, nil
	}
	params := &openapi_wireless.UpdateSimParams{}
	if status := resolveValue("status", current, config); status != "" {
		params.SetStatus(status)
	}
	if uniqueName := resolveValue("unique_name", current, config); uniqueName != "" {
		params.SetUniqueName(uniqueName)
	}
	sim, err := client.WirelessV1.UpdateSim(sid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(sim.Sid),
		"status": derefStr(sim.Status),
	}}, nil
}

// createRatePlanStep implements step.twilio_create_fleet (implemented as rate plan)
type createRatePlanStep struct {
	name       string
	moduleName string
}

func newCreateRatePlanStep(name string, config map[string]any) (*createRatePlanStep, error) {
	return &createRatePlanStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createRatePlanStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi_wireless.CreateRatePlanParams{}
	dataEnabled := resolveBool("data_enabled", current, config)
	params.SetDataEnabled(dataEnabled)
	if dataLimit := resolveInt("data_limit", current, config); dataLimit != 0 {
		params.SetDataLimit(dataLimit)
	}
	rp, err := client.WirelessV1.CreateRatePlan(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	dataEnabledResult := false
	if rp.DataEnabled != nil {
		dataEnabledResult = *rp.DataEnabled
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":          derefStr(rp.Sid),
		"data_enabled": dataEnabledResult,
	}}, nil
}

// sendCommandStep implements step.twilio_send_command
type sendCommandStep struct {
	name       string
	moduleName string
}

func newSendCommandStep(name string, config map[string]any) (*sendCommandStep, error) {
	return &sendCommandStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *sendCommandStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	command := resolveValue("command", current, config)
	if command == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "command is required"}}, nil
	}
	params := &openapi_wireless.CreateCommandParams{}
	params.SetCommand(command)
	if sim := resolveValue("sim", current, config); sim != "" {
		params.SetSim(sim)
	}
	cmd, err := client.WirelessV1.CreateCommand(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":     derefStr(cmd.Sid),
		"command": derefStr(cmd.Command),
		"status":  derefStr(cmd.Status),
	}}, nil
}
