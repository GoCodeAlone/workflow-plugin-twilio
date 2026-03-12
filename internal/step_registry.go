package internal

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// stepConstructor is a function that creates a StepInstance.
type stepConstructor func(name string, config map[string]any) (sdk.StepInstance, error)

// stepRegistry maps step type strings to constructor functions.
var stepRegistry = map[string]stepConstructor{
	// Messaging
	"step.twilio_send_sms":              func(n string, c map[string]any) (sdk.StepInstance, error) { return newSendSMSStep(n, c) },
	"step.twilio_send_mms":              func(n string, c map[string]any) (sdk.StepInstance, error) { return newSendMMSStep(n, c) },
	"step.twilio_send_whatsapp":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newSendWhatsappStep(n, c) },
	"step.twilio_list_messages":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newListMessagesStep(n, c) },
	"step.twilio_fetch_message":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchMessageStep(n, c) },
	"step.twilio_delete_message":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newDeleteMessageStep(n, c) },
	"step.twilio_fetch_media":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchMediaStep(n, c) },
	"step.twilio_create_messaging_service": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateMessagingServiceStep(n, c) },

	// Voice
	"step.twilio_create_call":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateCallStep(n, c) },
	"step.twilio_fetch_call":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchCallStep(n, c) },
	"step.twilio_list_calls":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newListCallsStep(n, c) },
	"step.twilio_update_call":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newUpdateCallStep(n, c) },
	"step.twilio_create_conference": func(n string, c map[string]any) (sdk.StepInstance, error) { return newListConferencesStep(n, c) },
	"step.twilio_list_conferences":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListConferencesStep(n, c) },
	"step.twilio_add_participant":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newAddParticipantStep(n, c) },
	"step.twilio_create_queue":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateQueueStep(n, c) },
	"step.twilio_fetch_recording":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchRecordingStep(n, c) },
	"step.twilio_list_recordings":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newListRecordingsStep(n, c) },
	"step.twilio_delete_recording":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newDeleteRecordingStep(n, c) },

	// Verify
	"step.twilio_send_verification":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newSendVerificationStep(n, c) },
	"step.twilio_check_verification": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCheckVerificationStep(n, c) },
	"step.twilio_create_verify_service": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateVerifyServiceStep(n, c) },
	"step.twilio_list_verify_services":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListVerifyServicesStep(n, c) },

	// Lookup
	"step.twilio_lookup_phone": func(n string, c map[string]any) (sdk.StepInstance, error) { return newLookupPhoneStep(n, c) },

	// Conversations
	"step.twilio_create_conversation":          func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateConversationStep(n, c) },
	"step.twilio_send_conversation_message":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newSendConversationMessageStep(n, c) },
	"step.twilio_add_conversation_participant": func(n string, c map[string]any) (sdk.StepInstance, error) { return newAddConversationParticipantStep(n, c) },
	"step.twilio_list_conversations":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newListConversationsStep(n, c) },
	"step.twilio_fetch_conversation":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchConversationStep(n, c) },
	"step.twilio_list_conversation_messages":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newListConversationMessagesStep(n, c) },
	"step.twilio_create_conversation_user":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateConversationUserStep(n, c) },

	// Video
	"step.twilio_create_room":          func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateRoomStep(n, c) },
	"step.twilio_list_rooms":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newListRoomsStep(n, c) },
	"step.twilio_fetch_room":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchRoomStep(n, c) },
	"step.twilio_complete_room":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompleteRoomStep(n, c) },
	"step.twilio_list_room_recordings": func(n string, c map[string]any) (sdk.StepInstance, error) { return newListRoomRecordingsStep(n, c) },
	"step.twilio_create_composition":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateCompositionStep(n, c) },

	// Notify
	"step.twilio_send_notification":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newSendNotificationStep(n, c) },
	"step.twilio_create_binding":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateBindingStep(n, c) },
	"step.twilio_list_bindings":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newListBindingsStep(n, c) },
	"step.twilio_create_notify_service": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateNotifyServiceStep(n, c) },

	// TaskRouter
	"step.twilio_create_workspace":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateWorkspaceStep(n, c) },
	"step.twilio_create_task":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateTaskStep(n, c) },
	"step.twilio_create_worker":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateWorkerStep(n, c) },
	"step.twilio_create_task_queue": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateTaskQueueStep(n, c) },
	"step.twilio_create_tr_workflow": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateTRWorkflowStep(n, c) },
	"step.twilio_list_tasks":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newListTasksStep(n, c) },
	"step.twilio_update_task":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newUpdateTaskStep(n, c) },

	// Phone Numbers
	"step.twilio_search_available": func(n string, c map[string]any) (sdk.StepInstance, error) { return newSearchAvailableStep(n, c) },
	"step.twilio_buy_number":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newBuyNumberStep(n, c) },
	"step.twilio_list_numbers":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newListNumbersStep(n, c) },
	"step.twilio_update_number":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newUpdateNumberStep(n, c) },
	"step.twilio_release_number":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newReleaseNumberStep(n, c) },

	// Studio
	"step.twilio_trigger_flow":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newTriggerFlowStep(n, c) },
	"step.twilio_list_flows":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newListFlowsStep(n, c) },
	"step.twilio_fetch_execution": func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchExecutionStep(n, c) },

	// Serverless
	"step.twilio_create_serverless_service": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateServerlessServiceStep(n, c) },
	"step.twilio_create_function":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateFunctionStep(n, c) },
	"step.twilio_create_build":              func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateBuildStep(n, c) },
	"step.twilio_list_serverless_services":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListServerlessServicesStep(n, c) },

	// Intelligence
	"step.twilio_create_transcript": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateTranscriptStep(n, c) },
	"step.twilio_fetch_transcript":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchTranscriptStep(n, c) },
	"step.twilio_list_transcripts":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListTranscriptsStep(n, c) },

	// Flex
	"step.twilio_create_flex_flow":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateFlexFlowStep(n, c) },
	"step.twilio_create_web_channel": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateWebChannelStep(n, c) },
	"step.twilio_list_flex_flows":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newListFlexFlowsStep(n, c) },

	// Proxy
	"step.twilio_create_proxy_service":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateProxyServiceStep(n, c) },
	"step.twilio_create_session":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateSessionStep(n, c) },
	"step.twilio_add_proxy_participant": func(n string, c map[string]any) (sdk.StepInstance, error) { return newAddProxyParticipantStep(n, c) },

	// Sync
	"step.twilio_create_sync_service": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateSyncServiceStep(n, c) },
	"step.twilio_create_document":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateDocumentStep(n, c) },
	"step.twilio_update_document":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newUpdateDocumentStep(n, c) },
	"step.twilio_create_sync_map":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateSyncMapStep(n, c) },
	"step.twilio_create_sync_list":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateSyncListStep(n, c) },

	// Wireless
	"step.twilio_list_sims":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newListSimsStep(n, c) },
	"step.twilio_fetch_sim":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchSimStep(n, c) },
	"step.twilio_update_sim":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newUpdateSimStep(n, c) },
	"step.twilio_create_fleet":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateRatePlanStep(n, c) },
	"step.twilio_send_command":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newSendCommandStep(n, c) },

	// Pricing
	"step.twilio_fetch_pricing":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchPricingStep(n, c) },
	"step.twilio_list_usage_records":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListUsageRecordsStep(n, c) },

	// Accounts
	"step.twilio_list_accounts":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListAccountsStep(n, c) },
	"step.twilio_create_api_key": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateApiKeyStep(n, c) },
	"step.twilio_list_api_keys":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListApiKeysStep(n, c) },

	// Content
	"step.twilio_create_content_template": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateContentTemplateStep(n, c) },
	"step.twilio_list_content_templates":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListContentTemplatesStep(n, c) },
	"step.twilio_fetch_content_template":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchContentTemplateStep(n, c) },

	// TrustHub
	"step.twilio_create_trust_product": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateTrustProductStep(n, c) },
	"step.twilio_list_trust_products":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newListTrustProductsStep(n, c) },
	"step.twilio_fetch_trust_product":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newFetchTrustProductStep(n, c) },

	// Assistants
	"step.twilio_create_assistant":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateAssistantStep(n, c) },
	"step.twilio_list_assistants":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newListAssistantsStep(n, c) },
	"step.twilio_create_knowledge_base": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCreateKnowledgeBaseStep(n, c) },
}

// createStep dispatches to the appropriate step constructor.
func createStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	constructor, ok := stepRegistry[typeName]
	if !ok {
		return nil, fmt.Errorf("twilio plugin: unknown step type %q", typeName)
	}
	return constructor(name, config)
}

// allStepTypes returns all registered step type strings.
func allStepTypes() []string {
	types := make([]string, 0, len(stepRegistry))
	for k := range stepRegistry {
		types = append(types, k)
	}
	return types
}
