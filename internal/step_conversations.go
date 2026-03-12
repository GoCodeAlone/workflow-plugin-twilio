package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_conv "github.com/twilio/twilio-go/rest/conversations/v1"
)

// createConversationStep implements step.twilio_create_conversation
type createConversationStep struct {
	name       string
	moduleName string
}

func newCreateConversationStep(name string, config map[string]any) (*createConversationStep, error) {
	return &createConversationStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createConversationStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi_conv.CreateConversationParams{}
	if fn := resolveValue("friendly_name", current, config); fn != "" {
		params.SetFriendlyName(fn)
	}
	conv, err := client.ConversationsV1.CreateConversation(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(conv.Sid),
		"friendly_name": derefStr(conv.FriendlyName),
		"state":         derefStr(conv.State),
	}}, nil
}

// sendConversationMessageStep implements step.twilio_send_conversation_message
type sendConversationMessageStep struct {
	name       string
	moduleName string
}

func newSendConversationMessageStep(name string, config map[string]any) (*sendConversationMessageStep, error) {
	return &sendConversationMessageStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *sendConversationMessageStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	conversationSid := resolveValue("conversation_sid", current, config)
	if conversationSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "conversation_sid is required"}}, nil
	}
	params := &openapi_conv.CreateConversationMessageParams{}
	if author := resolveValue("author", current, config); author != "" {
		params.SetAuthor(author)
	}
	if body := resolveValue("body", current, config); body != "" {
		params.SetBody(body)
	}
	msg, err := client.ConversationsV1.CreateConversationMessage(conversationSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":    derefStr(msg.Sid),
		"author": derefStr(msg.Author),
		"body":   derefStr(msg.Body),
		"index":  msg.Index,
	}}, nil
}

// addConversationParticipantStep implements step.twilio_add_conversation_participant
type addConversationParticipantStep struct {
	name       string
	moduleName string
}

func newAddConversationParticipantStep(name string, config map[string]any) (*addConversationParticipantStep, error) {
	return &addConversationParticipantStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *addConversationParticipantStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	conversationSid := resolveValue("conversation_sid", current, config)
	if conversationSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "conversation_sid is required"}}, nil
	}
	params := &openapi_conv.CreateConversationParticipantParams{}
	if identity := resolveValue("identity", current, config); identity != "" {
		params.SetIdentity(identity)
	}
	participant, err := client.ConversationsV1.CreateConversationParticipant(conversationSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":      derefStr(participant.Sid),
		"identity": derefStr(participant.Identity),
	}}, nil
}

// listConversationsStep implements step.twilio_list_conversations
type listConversationsStep struct {
	name       string
	moduleName string
}

func newListConversationsStep(name string, config map[string]any) (*listConversationsStep, error) {
	return &listConversationsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listConversationsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	convs, err := client.ConversationsV1.ListConversation(&openapi_conv.ListConversationParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(convs))
	for _, c := range convs {
		result = append(result, map[string]any{
			"sid":           derefStr(c.Sid),
			"friendly_name": derefStr(c.FriendlyName),
			"state":         derefStr(c.State),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"conversations": result, "count": len(result)}}, nil
}

// fetchConversationStep implements step.twilio_fetch_conversation
type fetchConversationStep struct {
	name       string
	moduleName string
}

func newFetchConversationStep(name string, config map[string]any) (*fetchConversationStep, error) {
	return &fetchConversationStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchConversationStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("conversation_sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "conversation_sid is required"}}, nil
	}
	conv, err := client.ConversationsV1.FetchConversation(sid)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(conv.Sid),
		"friendly_name": derefStr(conv.FriendlyName),
		"state":         derefStr(conv.State),
	}}, nil
}

// listConversationMessagesStep implements step.twilio_list_conversation_messages
type listConversationMessagesStep struct {
	name       string
	moduleName string
}

func newListConversationMessagesStep(name string, config map[string]any) (*listConversationMessagesStep, error) {
	return &listConversationMessagesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listConversationMessagesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	conversationSid := resolveValue("conversation_sid", current, config)
	if conversationSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "conversation_sid is required"}}, nil
	}
	msgs, err := client.ConversationsV1.ListConversationMessage(conversationSid, &openapi_conv.ListConversationMessageParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(msgs))
	for _, m := range msgs {
		result = append(result, map[string]any{
			"sid":    derefStr(m.Sid),
			"author": derefStr(m.Author),
			"body":   derefStr(m.Body),
			"index":  m.Index,
		})
	}
	return &sdk.StepResult{Output: map[string]any{"messages": result, "count": len(result)}}, nil
}

// createConversationUserStep implements step.twilio_create_conversation_user
type createConversationUserStep struct {
	name       string
	moduleName string
}

func newCreateConversationUserStep(name string, config map[string]any) (*createConversationUserStep, error) {
	return &createConversationUserStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createConversationUserStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	identity := resolveValue("identity", current, config)
	if identity == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "identity is required"}}, nil
	}
	params := &openapi_conv.CreateUserParams{}
	params.SetIdentity(identity)
	if fn := resolveValue("friendly_name", current, config); fn != "" {
		params.SetFriendlyName(fn)
	}
	user, err := client.ConversationsV1.CreateUser(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(user.Sid),
		"identity":      derefStr(user.Identity),
		"friendly_name": derefStr(user.FriendlyName),
	}}, nil
}
