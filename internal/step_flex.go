package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_flex "github.com/twilio/twilio-go/rest/flex/v1"
)

// createFlexFlowStep implements step.twilio_create_flex_flow
type createFlexFlowStep struct {
	name       string
	moduleName string
}

func newCreateFlexFlowStep(name string, config map[string]any) (*createFlexFlowStep, error) {
	return &createFlexFlowStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createFlexFlowStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	channelType := resolveValue("channel_type", current, config)
	if channelType == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "channel_type is required"}}, nil
	}
	params := &openapi_flex.CreateFlexFlowParams{}
	params.SetFriendlyName(friendlyName)
	params.SetChannelType(channelType)
	if chatServiceSid := resolveValue("chat_service_sid", current, config); chatServiceSid != "" {
		params.SetChatServiceSid(chatServiceSid)
	}
	flow, err := client.FlexV1.CreateFlexFlow(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(flow.Sid),
		"friendly_name": derefStr(flow.FriendlyName),
		"channel_type":  derefStr(flow.ChannelType),
	}}, nil
}

// createWebChannelStep implements step.twilio_create_web_channel
type createWebChannelStep struct {
	name       string
	moduleName string
}

func newCreateWebChannelStep(name string, config map[string]any) (*createWebChannelStep, error) {
	return &createWebChannelStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createWebChannelStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	flexFlowSid := resolveValue("flex_flow_sid", current, config)
	if flexFlowSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "flex_flow_sid is required"}}, nil
	}
	params := &openapi_flex.CreateWebChannelParams{}
	params.SetFlexFlowSid(flexFlowSid)
	if identity := resolveValue("identity", current, config); identity != "" {
		params.SetIdentity(identity)
	}
	if customerFriendlyName := resolveValue("customer_friendly_name", current, config); customerFriendlyName != "" {
		params.SetCustomerFriendlyName(customerFriendlyName)
	}
	ch, err := client.FlexV1.CreateWebChannel(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":          derefStr(ch.Sid),
		"flex_flow_sid": derefStr(ch.FlexFlowSid),
	}}, nil
}

// listFlexFlowsStep implements step.twilio_list_flex_flows
type listFlexFlowsStep struct {
	name       string
	moduleName string
}

func newListFlexFlowsStep(name string, config map[string]any) (*listFlexFlowsStep, error) {
	return &listFlexFlowsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listFlexFlowsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	flows, err := client.FlexV1.ListFlexFlow(&openapi_flex.ListFlexFlowParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(flows))
	for _, f := range flows {
		result = append(result, map[string]any{
			"sid":           derefStr(f.Sid),
			"friendly_name": derefStr(f.FriendlyName),
			"channel_type":  derefStr(f.ChannelType),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"flows": result, "count": len(result)}}, nil
}
