package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_sync "github.com/twilio/twilio-go/rest/sync/v1"
)

// createSyncServiceStep implements step.twilio_create_sync_service
type createSyncServiceStep struct {
	name       string
	moduleName string
}

func newCreateSyncServiceStep(name string, config map[string]any) (*createSyncServiceStep, error) {
	return &createSyncServiceStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createSyncServiceStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	params := &openapi_sync.CreateServiceParams{}
	if fn := resolveValue("friendly_name", current, config); fn != "" {
		params.SetFriendlyName(fn)
	}
	svc, err := client.SyncV1.CreateService(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(svc.Sid),
		"friendly_name": derefStr(svc.FriendlyName),
	}}, nil
}

// createDocumentStep implements step.twilio_create_document
type createDocumentStep struct {
	name       string
	moduleName string
}

func newCreateDocumentStep(name string, config map[string]any) (*createDocumentStep, error) {
	return &createDocumentStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createDocumentStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	params := &openapi_sync.CreateDocumentParams{}
	if un := resolveValue("unique_name", current, config); un != "" {
		params.SetUniqueName(un)
	}
	if data := resolveMap("data", current, config); data != nil {
		params.SetData(interface{}(data))
	}
	doc, err := client.SyncV1.CreateDocument(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(doc.Sid),
		"unique_name": derefStr(doc.UniqueName),
		"revision":    derefStr(doc.Revision),
	}}, nil
}

// updateDocumentStep implements step.twilio_update_document
type updateDocumentStep struct {
	name       string
	moduleName string
}

func newUpdateDocumentStep(name string, config map[string]any) (*updateDocumentStep, error) {
	return &updateDocumentStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *updateDocumentStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	documentSid := resolveValue("document_sid", current, config)
	if documentSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "document_sid is required"}}, nil
	}
	params := &openapi_sync.UpdateDocumentParams{}
	if data := resolveMap("data", current, config); data != nil {
		params.SetData(interface{}(data))
	}
	doc, err := client.SyncV1.UpdateDocument(serviceSid, documentSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":      derefStr(doc.Sid),
		"revision": derefStr(doc.Revision),
	}}, nil
}

// createSyncMapStep implements step.twilio_create_sync_map
type createSyncMapStep struct {
	name       string
	moduleName string
}

func newCreateSyncMapStep(name string, config map[string]any) (*createSyncMapStep, error) {
	return &createSyncMapStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createSyncMapStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	params := &openapi_sync.CreateSyncMapParams{}
	if un := resolveValue("unique_name", current, config); un != "" {
		params.SetUniqueName(un)
	}
	m, err := client.SyncV1.CreateSyncMap(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(m.Sid),
		"unique_name": derefStr(m.UniqueName),
	}}, nil
}

// createSyncListStep implements step.twilio_create_sync_list
type createSyncListStep struct {
	name       string
	moduleName string
}

func newCreateSyncListStep(name string, config map[string]any) (*createSyncListStep, error) {
	return &createSyncListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createSyncListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	serviceSid := resolveValue("service_sid", current, config)
	if serviceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "service_sid is required"}}, nil
	}
	params := &openapi_sync.CreateSyncListParams{}
	if un := resolveValue("unique_name", current, config); un != "" {
		params.SetUniqueName(un)
	}
	list, err := client.SyncV1.CreateSyncList(serviceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":         derefStr(list.Sid),
		"unique_name": derefStr(list.UniqueName),
	}}, nil
}
