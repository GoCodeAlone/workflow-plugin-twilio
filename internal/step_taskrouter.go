package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	openapi_tr "github.com/twilio/twilio-go/rest/taskrouter/v1"
)

// createWorkspaceStep implements step.twilio_create_workspace
type createWorkspaceStep struct {
	name       string
	moduleName string
}

func newCreateWorkspaceStep(name string, config map[string]any) (*createWorkspaceStep, error) {
	return &createWorkspaceStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createWorkspaceStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi_tr.CreateWorkspaceParams{}
	params.SetFriendlyName(friendlyName)
	ws, err := client.TaskrouterV1.CreateWorkspace(params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(ws.Sid),
		"friendly_name": derefStr(ws.FriendlyName),
	}}, nil
}

// createTaskStep implements step.twilio_create_task
type createTaskStep struct {
	name       string
	moduleName string
}

func newCreateTaskStep(name string, config map[string]any) (*createTaskStep, error) {
	return &createTaskStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createTaskStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	workspaceSid := resolveValue("workspace_sid", current, config)
	if workspaceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "workspace_sid is required"}}, nil
	}
	params := &openapi_tr.CreateTaskParams{}
	if taskChannel := resolveValue("task_channel", current, config); taskChannel != "" {
		params.SetTaskChannel(taskChannel)
	}
	if attributes := resolveValue("attributes", current, config); attributes != "" {
		params.SetAttributes(attributes)
	}
	task, err := client.TaskrouterV1.CreateTask(workspaceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":               derefStr(task.Sid),
		"task_queue_sid":    derefStr(task.TaskQueueSid),
		"assignment_status": derefStr(task.AssignmentStatus),
	}}, nil
}

// createWorkerStep implements step.twilio_create_worker
type createWorkerStep struct {
	name       string
	moduleName string
}

func newCreateWorkerStep(name string, config map[string]any) (*createWorkerStep, error) {
	return &createWorkerStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createWorkerStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	workspaceSid := resolveValue("workspace_sid", current, config)
	if workspaceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "workspace_sid is required"}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi_tr.CreateWorkerParams{}
	params.SetFriendlyName(friendlyName)
	if attributes := resolveValue("attributes", current, config); attributes != "" {
		params.SetAttributes(attributes)
	}
	worker, err := client.TaskrouterV1.CreateWorker(workspaceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(worker.Sid),
		"friendly_name": derefStr(worker.FriendlyName),
	}}, nil
}

// createTaskQueueStep implements step.twilio_create_task_queue
type createTaskQueueStep struct {
	name       string
	moduleName string
}

func newCreateTaskQueueStep(name string, config map[string]any) (*createTaskQueueStep, error) {
	return &createTaskQueueStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createTaskQueueStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	workspaceSid := resolveValue("workspace_sid", current, config)
	if workspaceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "workspace_sid is required"}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	params := &openapi_tr.CreateTaskQueueParams{}
	params.SetFriendlyName(friendlyName)
	if targetWorkers := resolveValue("target_workers", current, config); targetWorkers != "" {
		params.SetTargetWorkers(targetWorkers)
	}
	tq, err := client.TaskrouterV1.CreateTaskQueue(workspaceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(tq.Sid),
		"friendly_name": derefStr(tq.FriendlyName),
	}}, nil
}

// createTRWorkflowStep implements step.twilio_create_tr_workflow
type createTRWorkflowStep struct {
	name       string
	moduleName string
}

func newCreateTRWorkflowStep(name string, config map[string]any) (*createTRWorkflowStep, error) {
	return &createTRWorkflowStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *createTRWorkflowStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	workspaceSid := resolveValue("workspace_sid", current, config)
	if workspaceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "workspace_sid is required"}}, nil
	}
	friendlyName := resolveValue("friendly_name", current, config)
	if friendlyName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "friendly_name is required"}}, nil
	}
	configuration := resolveValue("configuration", current, config)
	if configuration == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "configuration is required"}}, nil
	}
	params := &openapi_tr.CreateWorkflowParams{}
	params.SetFriendlyName(friendlyName)
	params.SetConfiguration(configuration)
	wf, err := client.TaskrouterV1.CreateWorkflow(workspaceSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":           derefStr(wf.Sid),
		"friendly_name": derefStr(wf.FriendlyName),
	}}, nil
}

// listTasksStep implements step.twilio_list_tasks
type listTasksStep struct {
	name       string
	moduleName string
}

func newListTasksStep(name string, config map[string]any) (*listTasksStep, error) {
	return &listTasksStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *listTasksStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	workspaceSid := resolveValue("workspace_sid", current, config)
	if workspaceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "workspace_sid is required"}}, nil
	}
	tasks, err := client.TaskrouterV1.ListTask(workspaceSid, &openapi_tr.ListTaskParams{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	result := make([]map[string]any, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, map[string]any{
			"sid":               derefStr(t.Sid),
			"task_queue_sid":    derefStr(t.TaskQueueSid),
			"assignment_status": derefStr(t.AssignmentStatus),
		})
	}
	return &sdk.StepResult{Output: map[string]any{"tasks": result, "count": len(result)}}, nil
}

// updateTaskStep implements step.twilio_update_task
type updateTaskStep struct {
	name       string
	moduleName string
}

func newUpdateTaskStep(name string, config map[string]any) (*updateTaskStep, error) {
	return &updateTaskStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *updateTaskStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "twilio client not found: " + s.moduleName}}, nil
	}
	workspaceSid := resolveValue("workspace_sid", current, config)
	if workspaceSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "workspace_sid is required"}}, nil
	}
	taskSid := resolveValue("task_sid", current, config)
	if taskSid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "task_sid is required"}}, nil
	}
	params := &openapi_tr.UpdateTaskParams{}
	if status := resolveValue("assignment_status", current, config); status != "" {
		params.SetAssignmentStatus(status)
	}
	if reason := resolveValue("reason", current, config); reason != "" {
		params.SetReason(reason)
	}
	task, err := client.TaskrouterV1.UpdateTask(workspaceSid, taskSid, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{
		"sid":               derefStr(task.Sid),
		"assignment_status": derefStr(task.AssignmentStatus),
	}}, nil
}
