package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_content "github.com/twilio/twilio-go/rest/content/v1"
)

// createContentTemplateStep implements step.twilio_create_content_template
type createContentTemplateStep struct {
	name       string
	moduleName string
}

func newCreateContentTemplateStep(name string, config map[string]any) (*createContentTemplateStep, error) {
	return &createContentTemplateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createContentTemplateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	language := resolveValue("language", current, config)
	if language == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "language is required"}}, nil
	}
	params := &openapi_content.CreateContentParams{}
	params.SetContentCreateRequest(openapi_content.ContentCreateRequest{
		FriendlyName: friendlyName,
		Language:     language,
	})
	content, err := client.ContentV1.CreateContent(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(content.Sid),
		"friendly_name": derefStr(content.FriendlyName),
		"language":      derefStr(content.Language),
	}}, nil
}

// listContentTemplatesStep implements step.twilio_list_content_templates
type listContentTemplatesStep struct {
	name       string
	moduleName string
}

func newListContentTemplatesStep(name string, config map[string]any) (*listContentTemplatesStep, error) {
	return &listContentTemplatesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listContentTemplatesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	contents, err := client.ContentV1.ListContent(&openapi_content.ListContentParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(contents))
	for _, c := range contents {
		result = append(result, map[string]any{
			"sid":           derefStr(c.Sid),
			"friendly_name": derefStr(c.FriendlyName),
			"language":      derefStr(c.Language),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"templates": result, "count": len(result)}}, nil
}

// fetchContentTemplateStep implements step.twilio_fetch_content_template
type fetchContentTemplateStep struct {
	name       string
	moduleName string
}

func newFetchContentTemplateStep(name string, config map[string]any) (*fetchContentTemplateStep, error) {
	return &fetchContentTemplateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fetchContentTemplateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	sid := resolveValue("sid", current, config)
	if sid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "sid is required"}}, nil
	}
	content, err := client.ContentV1.FetchContent(sid)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(content.Sid),
		"friendly_name": derefStr(content.FriendlyName),
		"language":      derefStr(content.Language),
	}}, nil
}
