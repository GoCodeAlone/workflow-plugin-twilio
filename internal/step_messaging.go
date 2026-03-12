package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	openapi_messaging "github.com/twilio/twilio-go/rest/messaging/v1"
)

// sendSMSStep implements step.twilio_send_sms
type sendSMSStep struct {
	name       string
	moduleName string
}

func newSendSMSStep(name string, config map[string]any) (*sendSMSStep, error) {
	return &sendSMSStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *sendSMSStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	to := resolveValue("to", current, config)
	if to == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "to is required"}}, nil
	}
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(resolveValue("from", current, config))
	params.SetBody(resolveValue("body", current, config))
	msg, err := client.Api.CreateMessage(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(msg.Sid),
		"status": derefStr(msg.Status),
		"to":     derefStr(msg.To),
		"from":   derefStr(msg.From),
		"body":   derefStr(msg.Body),
	}}, nil
}

// sendMMSStep implements step.twilio_send_mms
type sendMMSStep struct {
	name       string
	moduleName string
}

func newSendMMSStep(name string, config map[string]any) (*sendMMSStep, error) {
	return &sendMMSStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *sendMMSStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	to := resolveValue("to", current, config)
	if to == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "to is required"}}, nil
	}
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(resolveValue("from", current, config))
	params.SetBody(resolveValue("body", current, config))
	if mediaUrl := resolveValue("media_url", current, config); mediaUrl != "" {
		params.SetMediaUrl([]string{mediaUrl})
	}
	msg, err := client.Api.CreateMessage(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(msg.Sid),
		"status": derefStr(msg.Status),
		"to":     derefStr(msg.To),
		"from":   derefStr(msg.From),
	}}, nil
}

// sendWhatsappStep implements step.twilio_send_whatsapp
type sendWhatsappStep struct {
	name       string
	moduleName string
}

func newSendWhatsappStep(name string, config map[string]any) (*sendWhatsappStep, error) {
	return &sendWhatsappStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *sendWhatsappStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	to := resolveValue("to", current, config)
	if to == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "to is required"}}, nil
	}
	from := resolveValue("from", current, config)
	// Ensure whatsapp: prefix
	if len(to) < 9 || to[:9] != "whatsapp:" {
		to = "whatsapp:" + to
	}
	if len(from) >= 9 && from[:9] != "whatsapp:" {
		from = "whatsapp:" + from
	} else if len(from) > 0 && len(from) < 9 {
		from = "whatsapp:" + from
	}
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(resolveValue("body", current, config))
	msg, err := client.Api.CreateMessage(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(msg.Sid),
		"status": derefStr(msg.Status),
	}}, nil
}

// listMessagesStep implements step.twilio_list_messages
type listMessagesStep struct {
	name       string
	moduleName string
}

func newListMessagesStep(name string, config map[string]any) (*listMessagesStep, error) {
	return &listMessagesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listMessagesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi.ListMessageParams{}
	if to := resolveValue("to", current, config); to != "" {
		params.SetTo(to)
	}
	if from := resolveValue("from", current, config); from != "" {
		params.SetFrom(from)
	}
	msgs, err := client.Api.ListMessage(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	messages := make([]map[string]any, 0, len(msgs))
	for _, m := range msgs {
		messages = append(messages, map[string]any{
			"sid":    derefStr(m.Sid),
			"status": derefStr(m.Status),
			"to":     derefStr(m.To),
			"from":   derefStr(m.From),
			"body":   derefStr(m.Body),
		})
	}
	return &sdk.StepResult{Output: map[string]any{
		"messages": messages,
		"count":    len(messages),
	}}, nil
}

// fetchMessageStep implements step.twilio_fetch_message
type fetchMessageStep struct {
	name       string
	moduleName string
}

func newFetchMessageStep(name string, config map[string]any) (*fetchMessageStep, error) {
	return &fetchMessageStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchMessageStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("message_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "message_sid is required"}}, nil
	}
	msg, err := client.Api.FetchMessage(sid, &openapi.FetchMessageParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(msg.Sid),
		"status": derefStr(msg.Status),
		"to":     derefStr(msg.To),
		"from":   derefStr(msg.From),
		"body":   derefStr(msg.Body),
	}}, nil
}

// deleteMessageStep implements step.twilio_delete_message
type deleteMessageStep struct {
	name       string
	moduleName string
}

func newDeleteMessageStep(name string, config map[string]any) (*deleteMessageStep, error) {
	return &deleteMessageStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *deleteMessageStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("message_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "message_sid is required"}}, nil
	}
	err := client.Api.DeleteMessage(sid, &openapi.DeleteMessageParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true}}, nil
}

// fetchMediaStep implements step.twilio_fetch_media
type fetchMediaStep struct {
	name       string
	moduleName string
}

func newFetchMediaStep(name string, config map[string]any) (*fetchMediaStep, error) {
	return &fetchMediaStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchMediaStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	messageSid := resolveValue("message_sid", current, config)
	mediaSid := resolveValue("media_sid", current, config)
	if messageSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "message_sid is required"}}, nil
	}
	if mediaSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "media_sid is required"}}, nil
	}
	media, err := client.Api.FetchMedia(messageSid, mediaSid, &openapi.FetchMediaParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":          derefStr(media.Sid),
		"content_type": derefStr(media.ContentType),
		"uri":          derefStr(media.Uri),
	}}, nil
}

// createMessagingServiceStep implements step.twilio_create_messaging_service
type createMessagingServiceStep struct {
	name       string
	moduleName string
}

func newCreateMessagingServiceStep(name string, config map[string]any) (*createMessagingServiceStep, error) {
	return &createMessagingServiceStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createMessagingServiceStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi_messaging.CreateServiceParams{}
	params.SetFriendlyName(friendlyName)
	svc, err := client.MessagingV1.CreateService(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(svc.Sid),
		"friendly_name": derefStr(svc.FriendlyName),
	}}, nil
}
