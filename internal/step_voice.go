package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// createCallStep implements step.twilio_create_call
type createCallStep struct {
	name       string
	moduleName string
}

func newCreateCallStep(name string, config map[string]any) (*createCallStep, error) {
	return &createCallStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createCallStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	to := resolveValue("to", current, config)
	if to == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "to is required"}}, nil
	}
	params := &openapi.CreateCallParams{}
	params.SetTo(to)
	params.SetFrom(resolveValue("from", current, config))
	params.SetUrl(resolveValue("url", current, config))
	call, err := client.Api.CreateCall(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(call.Sid),
		"status": derefStr(call.Status),
		"to":     derefStr(call.To),
		"from":   derefStr(call.From),
	}}, nil
}

// fetchCallStep implements step.twilio_fetch_call
type fetchCallStep struct {
	name       string
	moduleName string
}

func newFetchCallStep(name string, config map[string]any) (*fetchCallStep, error) {
	return &fetchCallStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchCallStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("call_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "call_sid is required"}}, nil
	}
	call, err := client.Api.FetchCall(sid, &openapi.FetchCallParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":      derefStr(call.Sid),
		"status":   derefStr(call.Status),
		"duration": derefStr(call.Duration),
		"to":       derefStr(call.To),
		"from":     derefStr(call.From),
	}}, nil
}

// listCallsStep implements step.twilio_list_calls
type listCallsStep struct {
	name       string
	moduleName string
}

func newListCallsStep(name string, config map[string]any) (*listCallsStep, error) {
	return &listCallsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listCallsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi.ListCallParams{}
	if to := resolveValue("to", current, config); to != "" {
		params.SetTo(to)
	}
	if from := resolveValue("from", current, config); from != "" {
		params.SetFrom(from)
	}
	calls, err := client.Api.ListCall(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(calls))
	for _, c := range calls {
		result = append(result, map[string]any{
			"sid":    derefStr(c.Sid),
			"status": derefStr(c.Status),
			"to":     derefStr(c.To),
			"from":   derefStr(c.From),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"calls": result, "count": len(result)}}, nil
}

// updateCallStep implements step.twilio_update_call
type updateCallStep struct {
	name       string
	moduleName string
}

func newUpdateCallStep(name string, config map[string]any) (*updateCallStep, error) {
	return &updateCallStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *updateCallStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("call_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "call_sid is required"}}, nil
	}
	params := &openapi.UpdateCallParams{}
	if status := resolveValue("status", current, config); status != "" {
		params.SetStatus(status)
	}
	if url := resolveValue("url", current, config); url != "" {
		params.SetUrl(url)
	}
	call, err := client.Api.UpdateCall(sid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(call.Sid),
		"status": derefStr(call.Status),
	}}, nil
}

// listConferencesStep implements step.twilio_list_conferences and step.twilio_create_conference
type listConferencesStep struct {
	name       string
	moduleName string
}

func newListConferencesStep(name string, config map[string]any) (*listConferencesStep, error) {
	return &listConferencesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listConferencesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi.ListConferenceParams{}
	conferences, err := client.Api.ListConference(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(conferences))
	for _, c := range conferences {
		result = append(result, map[string]any{
			"sid":           derefStr(c.Sid),
			"friendly_name": derefStr(c.FriendlyName),
			"status":        derefStr(c.Status),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"conferences": result, "count": len(result)}}, nil
}

// addParticipantStep implements step.twilio_add_participant
type addParticipantStep struct {
	name       string
	moduleName string
}

func newAddParticipantStep(name string, config map[string]any) (*addParticipantStep, error) {
	return &addParticipantStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *addParticipantStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	conferenceSid := resolveValue("conference_sid", current, config)
	if conferenceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "conference_sid is required"}}, nil
	}
	params := &openapi.CreateParticipantParams{}
	params.SetFrom(resolveValue("from", current, config))
	params.SetTo(resolveValue("to", current, config))
	if url := resolveValue("url", current, config); url != "" {
		params.SetStatusCallback(url)
	}
	participant, err := client.Api.CreateParticipant(conferenceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"call_sid":       derefStr(participant.CallSid),
		"conference_sid": derefStr(participant.ConferenceSid),
	}}, nil
}

// createQueueStep implements step.twilio_create_queue
type createQueueStep struct {
	name       string
	moduleName string
}

func newCreateQueueStep(name string, config map[string]any) (*createQueueStep, error) {
	return &createQueueStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createQueueStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi.CreateQueueParams{}
	params.SetFriendlyName(friendlyName)
	queue, err := client.Api.CreateQueue(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":          derefStr(queue.Sid),
		"friendly_name": derefStr(queue.FriendlyName),
		"current_size": queue.CurrentSize,
	}}, nil
}

// fetchRecordingStep implements step.twilio_fetch_recording
type fetchRecordingStep struct {
	name       string
	moduleName string
}

func newFetchRecordingStep(name string, config map[string]any) (*fetchRecordingStep, error) {
	return &fetchRecordingStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchRecordingStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("recording_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "recording_sid is required"}}, nil
	}
	rec, err := client.Api.FetchRecording(sid, &openapi.FetchRecordingParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":      derefStr(rec.Sid),
		"status":   derefStr(rec.Status),
		"duration": derefStr(rec.Duration),
		"call_sid": derefStr(rec.CallSid),
	}}, nil
}

// listRecordingsStep implements step.twilio_list_recordings
type listRecordingsStep struct {
	name       string
	moduleName string
}

func newListRecordingsStep(name string, config map[string]any) (*listRecordingsStep, error) {
	return &listRecordingsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listRecordingsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi.ListRecordingParams{}
	recordings, err := client.Api.ListRecording(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(recordings))
	for _, r := range recordings {
		result = append(result, map[string]any{
			"sid":      derefStr(r.Sid),
			"status":   derefStr(r.Status),
			"duration": derefStr(r.Duration),
			"call_sid": derefStr(r.CallSid),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"recordings": result, "count": len(result)}}, nil
}

// deleteRecordingStep implements step.twilio_delete_recording
type deleteRecordingStep struct {
	name       string
	moduleName string
}

func newDeleteRecordingStep(name string, config map[string]any) (*deleteRecordingStep, error) {
	return &deleteRecordingStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *deleteRecordingStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("recording_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "recording_sid is required"}}, nil
	}
	err := client.Api.DeleteRecording(sid, &openapi.DeleteRecordingParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true}}, nil
}
