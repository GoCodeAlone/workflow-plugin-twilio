package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_assistants "github.com/twilio/twilio-go/rest/assistants/v1"
)

// createAssistantStep implements step.twilio_create_assistant
type createAssistantStep struct {
	name       string
	moduleName string
}

func newCreateAssistantStep(name string, config map[string]any) (*createAssistantStep, error) {
	return &createAssistantStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createAssistantStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	name := resolveValue("name", current, config)
	if name == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name is required"}}, nil
	}
	params := &openapi_assistants.CreateAssistantParams{}
	params.SetAssistantsV1CreateAssistantRequest(openapi_assistants.AssistantsV1CreateAssistantRequest{
		Name: name,
	})
	assistant, err := client.AssistantsV1.CreateAssistant(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"id":   assistant.Id,
		"name": assistant.Name,
	}}, nil
}

// listAssistantsStep implements step.twilio_list_assistants
type listAssistantsStep struct {
	name       string
	moduleName string
}

func newListAssistantsStep(name string, config map[string]any) (*listAssistantsStep, error) {
	return &listAssistantsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listAssistantsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	assistants, err := client.AssistantsV1.ListAssistants(&openapi_assistants.ListAssistantsParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(assistants))
	for _, a := range assistants {
		result = append(result, map[string]any{
			"id":   a.Id,
			"name": a.Name,
		})
	}
	return &sdk.StepResult{Output: map[string]any{"assistants": result, "count": len(result)}}, nil
}

// createKnowledgeBaseStep implements step.twilio_create_knowledge_base
type createKnowledgeBaseStep struct {
	name       string
	moduleName string
}

func newCreateKnowledgeBaseStep(name string, config map[string]any) (*createKnowledgeBaseStep, error) {
	return &createKnowledgeBaseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createKnowledgeBaseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi_assistants.CreateKnowledgeParams{}
	knowledge, err := client.AssistantsV1.CreateKnowledge(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"id":   knowledge.Id,
		"name": knowledge.Name,
	}}, nil
}
