package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_video "github.com/twilio/twilio-go/rest/video/v1"
)

// createRoomStep implements step.twilio_create_room
type createRoomStep struct {
	name       string
	moduleName string
}

func newCreateRoomStep(name string, config map[string]any) (*createRoomStep, error) {
	return &createRoomStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createRoomStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi_video.CreateRoomParams{}
	if name := resolveValue("unique_name", current, config); name != "" {
		params.SetUniqueName(name)
	}
	if roomType := resolveValue("type", current, config); roomType != "" {
		params.SetType(roomType)
	}
	room, err := client.VideoV1.CreateRoom(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(room.Sid),
		"unique_name": derefStr(room.UniqueName),
		"status":      derefStr(room.Status),
		"type":        derefStr(room.Type),
	}}, nil
}

// listRoomsStep implements step.twilio_list_rooms
type listRoomsStep struct {
	name       string
	moduleName string
}

func newListRoomsStep(name string, config map[string]any) (*listRoomsStep, error) {
	return &listRoomsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listRoomsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi_video.ListRoomParams{}
	if status := resolveValue("status", current, config); status != "" {
		params.SetStatus(status)
	}
	rooms, err := client.VideoV1.ListRoom(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(rooms))
	for _, r := range rooms {
		result = append(result, map[string]any{
			"sid":         derefStr(r.Sid),
			"unique_name": derefStr(r.UniqueName),
			"status":      derefStr(r.Status),
			"type":        derefStr(r.Type),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"rooms": result, "count": len(result)}}, nil
}

// fetchRoomStep implements step.twilio_fetch_room
type fetchRoomStep struct {
	name       string
	moduleName string
}

func newFetchRoomStep(name string, config map[string]any) (*fetchRoomStep, error) {
	return &fetchRoomStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchRoomStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("room_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "room_sid is required"}}, nil
	}
	room, err := client.VideoV1.FetchRoom(sid)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	duration := 0
	if room.Duration != nil {
		duration = *room.Duration
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(room.Sid),
		"unique_name": derefStr(room.UniqueName),
		"status":      derefStr(room.Status),
		"type":        derefStr(room.Type),
		"duration":    duration,
	}}, nil
}

// completeRoomStep implements step.twilio_complete_room
type completeRoomStep struct {
	name       string
	moduleName string
}

func newCompleteRoomStep(name string, config map[string]any) (*completeRoomStep, error) {
	return &completeRoomStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *completeRoomStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("room_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "room_sid is required"}}, nil
	}
	params := &openapi_video.UpdateRoomParams{}
	params.SetStatus("completed")
	room, err := client.VideoV1.UpdateRoom(sid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(room.Sid),
		"status": derefStr(room.Status),
	}}, nil
}

// listRoomRecordingsStep implements step.twilio_list_room_recordings
type listRoomRecordingsStep struct {
	name       string
	moduleName string
}

func newListRoomRecordingsStep(name string, config map[string]any) (*listRoomRecordingsStep, error) {
	return &listRoomRecordingsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listRoomRecordingsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	roomSid := resolveValue("room_sid", current, config)
	if roomSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "room_sid is required"}}, nil
	}
	recordings, err := client.VideoV1.ListRoomRecording(roomSid, &openapi_video.ListRoomRecordingParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(recordings))
	for _, r := range recordings {
		result = append(result, map[string]any{
			"sid":      derefStr(r.Sid),
			"status":   derefStr(r.Status),
			"room_sid": derefStr(r.RoomSid),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"recordings": result, "count": len(result)}}, nil
}

// createCompositionStep implements step.twilio_create_composition
type createCompositionStep struct {
	name       string
	moduleName string
}

func newCreateCompositionStep(name string, config map[string]any) (*createCompositionStep, error) {
	return &createCompositionStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createCompositionStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	roomSid := resolveValue("room_sid", current, config)
	if roomSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "room_sid is required"}}, nil
	}
	params := &openapi_video.CreateCompositionParams{}
	params.SetRoomSid(roomSid)
	comp, err := client.VideoV1.CreateComposition(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":      derefStr(comp.Sid),
		"status":   derefStr(comp.Status),
		"room_sid": derefStr(comp.RoomSid),
	}}, nil
}
