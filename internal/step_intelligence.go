package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_intel "github.com/twilio/twilio-go/rest/intelligence/v2"
)

// createTranscriptStep implements step.twilio_create_transcript
type createTranscriptStep struct {
	name       string
	moduleName string
}

func newCreateTranscriptStep(name string, config map[string]any) (*createTranscriptStep, error) {
	return &createTranscriptStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createTranscriptStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	params := &openapi_intel.CreateTranscriptParams{}
	params.SetServiceSid(serviceSid)
	if channel := resolveMap("channel", current, config); channel != nil {
		params.SetChannel(interface{}(channel))
	}
	transcript, err := client.IntelligenceV2.CreateTranscript(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(transcript.Sid),
		"status":      derefStr(transcript.Status),
		"service_sid": derefStr(transcript.ServiceSid),
	}}, nil
}

// fetchTranscriptStep implements step.twilio_fetch_transcript
type fetchTranscriptStep struct {
	name       string
	moduleName string
}

func newFetchTranscriptStep(name string, config map[string]any) (*fetchTranscriptStep, error) {
	return &fetchTranscriptStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchTranscriptStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "sid is required"}}, nil
	}
	transcript, err := client.IntelligenceV2.FetchTranscript(sid)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	dateCreated := ""
	if transcript.DateCreated != nil {
		dateCreated = transcript.DateCreated.String()
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":          derefStr(transcript.Sid),
		"status":       derefStr(transcript.Status),
		"date_created": dateCreated,
	}}, nil
}

// listTranscriptsStep implements step.twilio_list_transcripts
type listTranscriptsStep struct {
	name       string
	moduleName string
}

func newListTranscriptsStep(name string, config map[string]any) (*listTranscriptsStep, error) {
	return &listTranscriptsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listTranscriptsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi_intel.ListTranscriptParams{}
	if serviceSid := resolveValue("service_sid", current, config); serviceSid != "" {
		params.SetServiceSid(serviceSid)
	}
	transcripts, err := client.IntelligenceV2.ListTranscript(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(transcripts))
	for _, t := range transcripts {
		result = append(result, map[string]any{
			"sid":    derefStr(t.Sid),
			"status": derefStr(t.Status),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"transcripts": result, "count": len(result)}}, nil
}
