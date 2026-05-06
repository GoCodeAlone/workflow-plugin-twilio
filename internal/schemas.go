package internal

import (
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
	"github.com/GoCodeAlone/workflow/schema"
)

// AllStepSchemas is the exported accessor used by tooling (e.g. tools/gen-schemas)
// to enumerate the full set of strict step contract descriptors. It wraps the
// unexported allStepSchemas() function which is used internally by twilioPlugin
// at runtime. At runtime the engine reads step schemas from plugin.json's
// stepSchemas field (populated via tools/gen-schemas) rather than calling this
// function directly.
func AllStepSchemas() []*schema.StepSchema {
	return allStepSchemas()
}

// twilioModuleSchemas returns the UI schema descriptors for all module types
// provided by this plugin. It implements sdk.SchemaProvider.
func twilioModuleSchemas() []sdk.ModuleSchemaData {
	return []sdk.ModuleSchemaData{
		{
			Type:        "twilio.provider",
			Label:       "Twilio",
			Category:    "communications",
			Description: "Twilio REST API client. Provides authentication and connectivity to the Twilio platform.",
			ConfigFields: []sdk.ConfigField{
				{Name: "accountSid", Type: "string", Description: "Twilio Account SID (starts with AC)", Required: true},
				{Name: "authToken", Type: "string", Description: "Twilio Auth Token (used with accountSid for basic auth)", Required: false},
				{Name: "apiKey", Type: "string", Description: "Twilio API Key SID (alternative to authToken)", Required: false},
				{Name: "apiSecret", Type: "string", Description: "Twilio API Key Secret (required when apiKey is set)", Required: false},
				{Name: "baseURL", Type: "string", Description: "Override the Twilio API base URL (useful for testing with a mock server)", Required: false},
				{Name: "region", Type: "string", Description: "Twilio region (e.g. us1, ie1)", Required: false},
				{Name: "edge", Type: "string", Description: "Twilio edge location", Required: false},
				{Name: "optional", Type: "boolean", Description: "When true, missing credentials are silently ignored instead of returning an error", Required: false},
			},
		},
	}
}

// allStepSchemas returns strict contract descriptors for every step type
// advertised by this plugin. These are embedded in plugin.json as stepSchemas
// and also exposed at runtime via the gRPC GetStepSchemas RPC (future).
func allStepSchemas() []*schema.StepSchema {
	moduleField := schema.ConfigFieldDef{
		Key:         "module",
		Label:       "Module",
		Type:        schema.FieldTypeString,
		Description: "Name of the twilio.provider module instance to use (defaults to \"twilio\")",
		Required:    false,
	}

	return []*schema.StepSchema{
		// ── Messaging ────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_send_sms",
			Plugin:      "workflow-plugin-twilio",
			Description: "Send an SMS message via Twilio.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Destination phone number in E.164 format", Required: true},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Sender phone number or messaging service SID"},
				{Key: "body", Label: "Body", Type: schema.FieldTypeString, Description: "Message body text"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Twilio message SID"},
				{Key: "status", Type: "string", Description: "Message status (e.g. queued, sent, delivered)"},
				{Key: "to", Type: "string", Description: "Destination number"},
				{Key: "from", Type: "string", Description: "Sender number"},
				{Key: "body", Type: "string", Description: "Message body"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_send_mms",
			Plugin:      "workflow-plugin-twilio",
			Description: "Send an MMS message (with optional media URL) via Twilio.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Destination phone number in E.164 format", Required: true},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Sender phone number"},
				{Key: "body", Label: "Body", Type: schema.FieldTypeString, Description: "Message body text"},
				{Key: "media_url", Label: "Media URL", Type: schema.FieldTypeString, Description: "Public URL of the media attachment"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Twilio message SID"},
				{Key: "status", Type: "string", Description: "Message status"},
				{Key: "to", Type: "string", Description: "Destination number"},
				{Key: "from", Type: "string", Description: "Sender number"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_send_whatsapp",
			Plugin:      "workflow-plugin-twilio",
			Description: "Send a WhatsApp message via Twilio.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Destination WhatsApp number (whatsapp: prefix added automatically)", Required: true},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Sender WhatsApp number"},
				{Key: "body", Label: "Body", Type: schema.FieldTypeString, Description: "Message body text"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Twilio message SID"},
				{Key: "status", Type: "string", Description: "Message status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_messages",
			Plugin:      "workflow-plugin-twilio",
			Description: "List messages from the Twilio account, optionally filtered by to/from.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Filter by destination phone number"},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Filter by sender phone number"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "messages", Type: "array", Description: "List of message objects"},
				{Key: "count", Type: "number", Description: "Number of messages returned"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_message",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a single message by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "message_sid", Label: "Message SID", Type: schema.FieldTypeString, Description: "Twilio message SID (starts with SM)", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Message SID"},
				{Key: "status", Type: "string", Description: "Message status"},
				{Key: "to", Type: "string", Description: "Destination number"},
				{Key: "from", Type: "string", Description: "Sender number"},
				{Key: "body", Type: "string", Description: "Message body"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_delete_message",
			Plugin:      "workflow-plugin-twilio",
			Description: "Delete a message by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "message_sid", Label: "Message SID", Type: schema.FieldTypeString, Description: "Twilio message SID to delete", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "deleted", Type: "boolean", Description: "true when the message was deleted"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_media",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a media resource attached to a message.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "message_sid", Label: "Message SID", Type: schema.FieldTypeString, Description: "Parent message SID", Required: true},
				{Key: "media_sid", Label: "Media SID", Type: schema.FieldTypeString, Description: "Media resource SID", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Media SID"},
				{Key: "content_type", Type: "string", Description: "MIME content type"},
				{Key: "uri", Type: "string", Description: "Media URI"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_messaging_service",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Messaging Service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Human-readable name for the messaging service", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Messaging service SID"},
				{Key: "friendly_name", Type: "string", Description: "Service friendly name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Voice ─────────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_call",
			Plugin:      "workflow-plugin-twilio",
			Description: "Initiate an outbound phone call.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Destination phone number in E.164 format", Required: true},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Caller ID phone number"},
				{Key: "url", Label: "TwiML URL", Type: schema.FieldTypeString, Description: "URL that returns TwiML instructions for the call"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Call SID"},
				{Key: "status", Type: "string", Description: "Call status"},
				{Key: "to", Type: "string", Description: "Destination number"},
				{Key: "from", Type: "string", Description: "Caller ID"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_call",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch details of a call by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "call_sid", Label: "Call SID", Type: schema.FieldTypeString, Description: "Twilio call SID (starts with CA)", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Call SID"},
				{Key: "status", Type: "string", Description: "Call status"},
				{Key: "duration", Type: "string", Description: "Call duration in seconds"},
				{Key: "to", Type: "string", Description: "Destination number"},
				{Key: "from", Type: "string", Description: "Caller ID"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_calls",
			Plugin:      "workflow-plugin-twilio",
			Description: "List calls, optionally filtered by to/from.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Filter by destination number"},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Filter by caller number"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "calls", Type: "array", Description: "List of call objects"},
				{Key: "count", Type: "number", Description: "Number of calls returned"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_update_call",
			Plugin:      "workflow-plugin-twilio",
			Description: "Update an in-progress call (e.g. redirect or terminate).",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "call_sid", Label: "Call SID", Type: schema.FieldTypeString, Description: "SID of the call to update", Required: true},
				{Key: "status", Label: "Status", Type: schema.FieldTypeString, Description: "New call status (e.g. completed, canceled)"},
				{Key: "url", Label: "TwiML URL", Type: schema.FieldTypeString, Description: "New TwiML URL to redirect the call to"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Call SID"},
				{Key: "status", Type: "string", Description: "Updated call status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_conference",
			Plugin:      "workflow-plugin-twilio",
			Description: "List conferences (alias for list_conferences; conference creation is implicit via TwiML).",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "conferences", Type: "array", Description: "List of conference objects"},
				{Key: "count", Type: "number", Description: "Number of conferences"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_conferences",
			Plugin:      "workflow-plugin-twilio",
			Description: "List all conferences in the account.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "conferences", Type: "array", Description: "List of conference objects"},
				{Key: "count", Type: "number", Description: "Number of conferences"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_add_participant",
			Plugin:      "workflow-plugin-twilio",
			Description: "Add a PSTN participant to a conference.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "conference_sid", Label: "Conference SID", Type: schema.FieldTypeString, Description: "Conference SID to add participant to", Required: true},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Caller ID to use for the new leg"},
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Phone number of participant to dial"},
				{Key: "url", Label: "Status Callback URL", Type: schema.FieldTypeString, Description: "URL for participant status callbacks"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "call_sid", Type: "string", Description: "Call SID of the new participant leg"},
				{Key: "conference_sid", Type: "string", Description: "Conference SID"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_queue",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a call queue.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Queue name", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Queue SID"},
				{Key: "friendly_name", Type: "string", Description: "Queue name"},
				{Key: "current_size", Type: "number", Description: "Current number of calls in the queue"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_recording",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a call recording by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "recording_sid", Label: "Recording SID", Type: schema.FieldTypeString, Description: "Twilio recording SID (starts with RE)", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Recording SID"},
				{Key: "duration", Type: "string", Description: "Duration in seconds"},
				{Key: "status", Type: "string", Description: "Recording status"},
				{Key: "call_sid", Type: "string", Description: "Call SID associated with the recording"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_recordings",
			Plugin:      "workflow-plugin-twilio",
			Description: "List call recordings, optionally filtered by call SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "call_sid", Label: "Call SID", Type: schema.FieldTypeString, Description: "Filter by call SID"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "recordings", Type: "array", Description: "List of recording objects"},
				{Key: "count", Type: "number", Description: "Number of recordings"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_delete_recording",
			Plugin:      "workflow-plugin-twilio",
			Description: "Delete a call recording by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "recording_sid", Label: "Recording SID", Type: schema.FieldTypeString, Description: "Recording SID to delete", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "deleted", Type: "boolean", Description: "true when the recording was deleted"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Verify ───────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_send_verification",
			Plugin:      "workflow-plugin-twilio",
			Description: "Send a phone verification code via Twilio Verify.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Verify service SID", Required: true},
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Phone number to verify", Required: true},
				{Key: "channel", Label: "Channel", Type: schema.FieldTypeString, Description: "Delivery channel: sms, call, email, totp"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Verification SID"},
				{Key: "status", Type: "string", Description: "Verification status"},
				{Key: "to", Type: "string", Description: "Destination"},
				{Key: "channel", Type: "string", Description: "Channel used"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_check_verification",
			Plugin:      "workflow-plugin-twilio",
			Description: "Check a verification code submitted by the user.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Verify service SID", Required: true},
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Phone number being verified", Required: true},
				{Key: "code", Label: "Code", Type: schema.FieldTypeString, Description: "Verification code entered by the user", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "status", Type: "string", Description: "approved or pending"},
				{Key: "valid", Type: "boolean", Description: "true if the code was correct"},
				{Key: "to", Type: "string", Description: "Phone number that was verified"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_verify_service",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Verify Service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Human-readable name for the service", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Verify service SID"},
				{Key: "friendly_name", Type: "string", Description: "Service name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_verify_services",
			Plugin:      "workflow-plugin-twilio",
			Description: "List all Twilio Verify Services.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "services", Type: "array", Description: "List of verify service objects"},
				{Key: "count", Type: "number", Description: "Number of services"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Lookup ───────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_lookup_phone",
			Plugin:      "workflow-plugin-twilio",
			Description: "Look up phone number metadata using Twilio Lookup.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "phone_number", Label: "Phone Number", Type: schema.FieldTypeString, Description: "Phone number to look up in E.164 format", Required: true},
				{Key: "fields", Label: "Fields", Type: schema.FieldTypeString, Description: "Comma-separated list of lookup fields (e.g. line_type_intelligence, caller_name)"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "phone_number", Type: "string", Description: "E.164 formatted number"},
				{Key: "country_code", Type: "string", Description: "ISO country code"},
				{Key: "national_format", Type: "string", Description: "Locally formatted number"},
				{Key: "valid", Type: "boolean", Description: "Whether the number is valid"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Conversations ─────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_conversation",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a new Twilio Conversation.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Conversation display name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Conversation SID"},
				{Key: "friendly_name", Type: "string", Description: "Conversation name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_send_conversation_message",
			Plugin:      "workflow-plugin-twilio",
			Description: "Send a message to a Twilio Conversation.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "conversation_sid", Label: "Conversation SID", Type: schema.FieldTypeString, Description: "Target conversation SID", Required: true},
				{Key: "author", Label: "Author", Type: schema.FieldTypeString, Description: "Message author identity"},
				{Key: "body", Label: "Body", Type: schema.FieldTypeString, Description: "Message body text"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Message SID"},
				{Key: "body", Type: "string", Description: "Message body"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_add_conversation_participant",
			Plugin:      "workflow-plugin-twilio",
			Description: "Add a participant to a Twilio Conversation.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "conversation_sid", Label: "Conversation SID", Type: schema.FieldTypeString, Description: "Target conversation SID", Required: true},
				{Key: "identity", Label: "Identity", Type: schema.FieldTypeString, Description: "Participant identity (for chat participants)"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Participant SID"},
				{Key: "identity", Type: "string", Description: "Participant identity"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_conversations",
			Plugin:      "workflow-plugin-twilio",
			Description: "List all Conversations.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "conversations", Type: "array", Description: "List of conversation objects"},
				{Key: "count", Type: "number", Description: "Number of conversations"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_conversation",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a Conversation by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "conversation_sid", Label: "Conversation SID", Type: schema.FieldTypeString, Description: "Conversation SID to fetch", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Conversation SID"},
				{Key: "friendly_name", Type: "string", Description: "Conversation name"},
				{Key: "state", Type: "string", Description: "Conversation state"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_conversation_messages",
			Plugin:      "workflow-plugin-twilio",
			Description: "List messages in a Conversation.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "conversation_sid", Label: "Conversation SID", Type: schema.FieldTypeString, Description: "Conversation to list messages for", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "messages", Type: "array", Description: "List of message objects"},
				{Key: "count", Type: "number", Description: "Number of messages"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_conversation_user",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Conversations user.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "identity", Label: "Identity", Type: schema.FieldTypeString, Description: "Unique user identity", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Display name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "User SID"},
				{Key: "identity", Type: "string", Description: "User identity"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Video ─────────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_room",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Video Room.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "Room unique name"},
				{Key: "type", Label: "Type", Type: schema.FieldTypeString, Description: "Room type: go, peer-to-peer, group, group-small"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Room SID"},
				{Key: "unique_name", Type: "string", Description: "Room unique name"},
				{Key: "status", Type: "string", Description: "Room status"},
				{Key: "type", Type: "string", Description: "Room type"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_rooms",
			Plugin:      "workflow-plugin-twilio",
			Description: "List Video Rooms, optionally filtered by status.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "status", Label: "Status", Type: schema.FieldTypeString, Description: "Filter by status: in-progress or completed"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "rooms", Type: "array", Description: "List of room objects"},
				{Key: "count", Type: "number", Description: "Number of rooms"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_room",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a Video Room by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "room_sid", Label: "Room SID", Type: schema.FieldTypeString, Description: "Room SID to fetch", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Room SID"},
				{Key: "unique_name", Type: "string", Description: "Room unique name"},
				{Key: "status", Type: "string", Description: "Room status"},
				{Key: "type", Type: "string", Description: "Room type"},
				{Key: "duration", Type: "number", Description: "Room duration in seconds"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_complete_room",
			Plugin:      "workflow-plugin-twilio",
			Description: "End a Video Room by setting its status to completed.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "room_sid", Label: "Room SID", Type: schema.FieldTypeString, Description: "Room SID to complete", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Room SID"},
				{Key: "status", Type: "string", Description: "Updated room status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_room_recordings",
			Plugin:      "workflow-plugin-twilio",
			Description: "List recordings for a Video Room.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "room_sid", Label: "Room SID", Type: schema.FieldTypeString, Description: "Room SID to list recordings for", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "recordings", Type: "array", Description: "List of recording objects"},
				{Key: "count", Type: "number", Description: "Number of recordings"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_composition",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Video Room composition from recordings.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "room_sid", Label: "Room SID", Type: schema.FieldTypeString, Description: "Room SID to compose", Required: true},
				{Key: "audio_sources", Label: "Audio Sources", Type: schema.FieldTypeString, Description: "Comma-separated list of audio source track names"},
				{Key: "video_layout", Label: "Video Layout", Type: schema.FieldTypeString, Description: "JSON video layout definition"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Composition SID"},
				{Key: "status", Type: "string", Description: "Composition status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Notify ───────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_send_notification",
			Plugin:      "workflow-plugin-twilio",
			Description: "Send a push notification via Twilio Notify.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Notify service SID", Required: true},
				{Key: "body", Label: "Body", Type: schema.FieldTypeString, Description: "Notification body text"},
				{Key: "title", Label: "Title", Type: schema.FieldTypeString, Description: "Notification title"},
				{Key: "identity", Label: "Identity", Type: schema.FieldTypeString, Description: "Target user identity"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Notification SID"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_binding",
			Plugin:      "workflow-plugin-twilio",
			Description: "Register a device binding for push notifications.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Notify service SID", Required: true},
				{Key: "identity", Label: "Identity", Type: schema.FieldTypeString, Description: "User identity", Required: true},
				{Key: "binding_type", Label: "Binding Type", Type: schema.FieldTypeString, Description: "Binding type: apn, fcm, gcm, sms, facebook-messenger", Required: true},
				{Key: "address", Label: "Address", Type: schema.FieldTypeString, Description: "Device token or address", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Binding SID"},
				{Key: "identity", Type: "string", Description: "User identity"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_bindings",
			Plugin:      "workflow-plugin-twilio",
			Description: "List device bindings for a Notify service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Notify service SID", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "bindings", Type: "array", Description: "List of binding objects"},
				{Key: "count", Type: "number", Description: "Number of bindings"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_notify_service",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Notify Service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Service name", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Notify service SID"},
				{Key: "friendly_name", Type: "string", Description: "Service name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── TaskRouter ────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_workspace",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a TaskRouter Workspace.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Workspace name", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Workspace SID"},
				{Key: "friendly_name", Type: "string", Description: "Workspace name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_task",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a TaskRouter Task.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "workspace_sid", Label: "Workspace SID", Type: schema.FieldTypeString, Description: "Target workspace SID", Required: true},
				{Key: "task_channel", Label: "Task Channel", Type: schema.FieldTypeString, Description: "Task channel unique name (e.g. voice, chat)"},
				{Key: "attributes", Label: "Attributes", Type: schema.FieldTypeString, Description: "JSON string of task attributes"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Task SID"},
				{Key: "assignment_status", Type: "string", Description: "Task assignment status"},
				{Key: "task_queue_sid", Type: "string", Description: "Task queue SID the task was routed to"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_worker",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a TaskRouter Worker.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "workspace_sid", Label: "Workspace SID", Type: schema.FieldTypeString, Description: "Target workspace SID", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Worker name", Required: true},
				{Key: "attributes", Label: "Attributes", Type: schema.FieldTypeString, Description: "JSON string of worker attributes"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Worker SID"},
				{Key: "friendly_name", Type: "string", Description: "Worker name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_task_queue",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a TaskRouter Task Queue.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "workspace_sid", Label: "Workspace SID", Type: schema.FieldTypeString, Description: "Target workspace SID", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Queue name", Required: true},
				{Key: "target_workers", Label: "Target Workers", Type: schema.FieldTypeString, Description: "Worker expression for routing"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Task queue SID"},
				{Key: "friendly_name", Type: "string", Description: "Queue name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_tr_workflow",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a TaskRouter Workflow.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "workspace_sid", Label: "Workspace SID", Type: schema.FieldTypeString, Description: "Target workspace SID", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Workflow name", Required: true},
				{Key: "configuration", Label: "Configuration", Type: schema.FieldTypeString, Description: "JSON workflow configuration string"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Workflow SID"},
				{Key: "friendly_name", Type: "string", Description: "Workflow name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_tasks",
			Plugin:      "workflow-plugin-twilio",
			Description: "List tasks in a TaskRouter Workspace.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "workspace_sid", Label: "Workspace SID", Type: schema.FieldTypeString, Description: "Target workspace SID", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "tasks", Type: "array", Description: "List of task objects"},
				{Key: "count", Type: "number", Description: "Number of tasks"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_update_task",
			Plugin:      "workflow-plugin-twilio",
			Description: "Update a TaskRouter Task (e.g. reassign or complete).",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "workspace_sid", Label: "Workspace SID", Type: schema.FieldTypeString, Description: "Target workspace SID", Required: true},
				{Key: "task_sid", Label: "Task SID", Type: schema.FieldTypeString, Description: "Task SID to update", Required: true},
				{Key: "assignment_status", Label: "Assignment Status", Type: schema.FieldTypeString, Description: "New assignment status (e.g. completed, canceled)"},
				{Key: "reason", Label: "Reason", Type: schema.FieldTypeString, Description: "Reason for the status update"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Task SID"},
				{Key: "assignment_status", Type: "string", Description: "Updated assignment status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Phone Numbers ─────────────────────────────────────────────────────
		{
			Type:        "step.twilio_search_available",
			Plugin:      "workflow-plugin-twilio",
			Description: "Search for available phone numbers to purchase.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "country_code", Label: "Country Code", Type: schema.FieldTypeString, Description: "ISO country code (e.g. US, GB)", Required: true},
				{Key: "contains", Label: "Contains", Type: schema.FieldTypeString, Description: "Pattern to match in the phone number"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "numbers", Type: "array", Description: "List of available number objects"},
				{Key: "count", Type: "number", Description: "Number of results"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_buy_number",
			Plugin:      "workflow-plugin-twilio",
			Description: "Purchase a phone number.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "phone_number", Label: "Phone Number", Type: schema.FieldTypeString, Description: "E.164 phone number to purchase", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Label for the purchased number"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Incoming phone number SID"},
				{Key: "phone_number", Type: "string", Description: "Purchased E.164 phone number"},
				{Key: "friendly_name", Type: "string", Description: "Number label"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_numbers",
			Plugin:      "workflow-plugin-twilio",
			Description: "List incoming phone numbers in the account.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "numbers", Type: "array", Description: "List of phone number objects"},
				{Key: "count", Type: "number", Description: "Number of results"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_update_number",
			Plugin:      "workflow-plugin-twilio",
			Description: "Update an incoming phone number's configuration.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "sid", Label: "SID", Type: schema.FieldTypeString, Description: "Phone number SID", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "New label"},
				{Key: "sms_url", Label: "SMS URL", Type: schema.FieldTypeString, Description: "Webhook URL for incoming SMS"},
				{Key: "voice_url", Label: "Voice URL", Type: schema.FieldTypeString, Description: "Webhook URL for incoming calls"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Phone number SID"},
				{Key: "friendly_name", Type: "string", Description: "Updated label"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_release_number",
			Plugin:      "workflow-plugin-twilio",
			Description: "Release (delete) an incoming phone number.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "sid", Label: "SID", Type: schema.FieldTypeString, Description: "Phone number SID to release", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "deleted", Type: "boolean", Description: "true when the number was released"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Studio ───────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_trigger_flow",
			Plugin:      "workflow-plugin-twilio",
			Description: "Trigger a Twilio Studio Flow execution.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "flow_sid", Label: "Flow SID", Type: schema.FieldTypeString, Description: "Studio flow SID (starts with FW)", Required: true},
				{Key: "to", Label: "To", Type: schema.FieldTypeString, Description: "Contact phone number"},
				{Key: "from", Label: "From", Type: schema.FieldTypeString, Description: "Twilio phone number"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Execution SID"},
				{Key: "status", Type: "string", Description: "Execution status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_flows",
			Plugin:      "workflow-plugin-twilio",
			Description: "List Studio Flows.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "flows", Type: "array", Description: "List of flow objects"},
				{Key: "count", Type: "number", Description: "Number of flows"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_execution",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a Studio Flow execution by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "flow_sid", Label: "Flow SID", Type: schema.FieldTypeString, Description: "Studio flow SID", Required: true},
				{Key: "execution_sid", Label: "Execution SID", Type: schema.FieldTypeString, Description: "Execution SID to fetch", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Execution SID"},
				{Key: "status", Type: "string", Description: "Execution status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Serverless ────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_serverless_service",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Serverless Service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "Unique identifier for the service", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Human-readable service name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Serverless service SID"},
				{Key: "unique_name", Type: "string", Description: "Service unique name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_function",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Serverless Function within a Service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Parent serverless service SID", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Function name", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Function SID"},
				{Key: "friendly_name", Type: "string", Description: "Function name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_build",
			Plugin:      "workflow-plugin-twilio",
			Description: "Trigger a Serverless Build to deploy assets and functions.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Parent serverless service SID", Required: true},
				{Key: "dependencies", Label: "Dependencies", Type: schema.FieldTypeString, Description: "JSON-encoded dependencies list"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Build SID"},
				{Key: "status", Type: "string", Description: "Build status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_serverless_services",
			Plugin:      "workflow-plugin-twilio",
			Description: "List Serverless Services.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "services", Type: "array", Description: "List of serverless service objects"},
				{Key: "count", Type: "number", Description: "Number of services"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Intelligence ──────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_transcript",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create an Intelligence transcript from a call recording.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Intelligence service SID", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Transcript SID"},
				{Key: "status", Type: "string", Description: "Transcript status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_transcript",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a transcript by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "sid", Label: "SID", Type: schema.FieldTypeString, Description: "Transcript SID to fetch", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Transcript SID"},
				{Key: "status", Type: "string", Description: "Transcript status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_transcripts",
			Plugin:      "workflow-plugin-twilio",
			Description: "List transcripts, optionally filtered by service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Filter by intelligence service SID"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "transcripts", Type: "array", Description: "List of transcript objects"},
				{Key: "count", Type: "number", Description: "Number of transcripts"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Flex ──────────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_flex_flow",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Flex Flow.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Flow name", Required: true},
				{Key: "channel_type", Label: "Channel Type", Type: schema.FieldTypeString, Description: "Channel type: web, sms, facebook, whatsapp, line, custom", Required: true},
				{Key: "chat_service_sid", Label: "Chat Service SID", Type: schema.FieldTypeString, Description: "Flex chat service SID"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Flex flow SID"},
				{Key: "friendly_name", Type: "string", Description: "Flow name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_web_channel",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Flex Web Channel (chat session).",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "flex_flow_sid", Label: "Flex Flow SID", Type: schema.FieldTypeString, Description: "Flex flow to use", Required: true},
				{Key: "identity", Label: "Identity", Type: schema.FieldTypeString, Description: "Customer identity"},
				{Key: "customer_friendly_name", Label: "Customer Name", Type: schema.FieldTypeString, Description: "Customer display name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Web channel SID"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_flex_flows",
			Plugin:      "workflow-plugin-twilio",
			Description: "List Flex Flows.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "flows", Type: "array", Description: "List of flex flow objects"},
				{Key: "count", Type: "number", Description: "Number of flows"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Proxy ─────────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_proxy_service",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Proxy Service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "Unique name for the service", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Proxy service SID"},
				{Key: "unique_name", Type: "string", Description: "Service unique name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_session",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Proxy Session.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Proxy service SID", Required: true},
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "Optional session unique name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Session SID"},
				{Key: "unique_name", Type: "string", Description: "Session unique name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_add_proxy_participant",
			Plugin:      "workflow-plugin-twilio",
			Description: "Add a participant to a Proxy Session.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Proxy service SID", Required: true},
				{Key: "session_sid", Label: "Session SID", Type: schema.FieldTypeString, Description: "Session SID to add participant to", Required: true},
				{Key: "identifier", Label: "Identifier", Type: schema.FieldTypeString, Description: "Phone number or identifier for the participant", Required: true},
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Participant display name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Participant SID"},
				{Key: "identifier", Type: "string", Description: "Participant identifier"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Sync ──────────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_sync_service",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Sync Service.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Service name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Sync service SID"},
				{Key: "friendly_name", Type: "string", Description: "Service name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_document",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Sync Document.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Sync service SID", Required: true},
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "Document unique name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Document SID"},
				{Key: "unique_name", Type: "string", Description: "Document unique name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_update_document",
			Plugin:      "workflow-plugin-twilio",
			Description: "Update a Sync Document's data.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Sync service SID", Required: true},
				{Key: "document_sid", Label: "Document SID", Type: schema.FieldTypeString, Description: "Document SID or unique name", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Document SID"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_sync_map",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Sync Map.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Sync service SID", Required: true},
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "Map unique name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Sync map SID"},
				{Key: "unique_name", Type: "string", Description: "Map unique name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_sync_list",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Sync List.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "service_sid", Label: "Service SID", Type: schema.FieldTypeString, Description: "Sync service SID", Required: true},
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "List unique name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Sync list SID"},
				{Key: "unique_name", Type: "string", Description: "List unique name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Wireless ──────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_list_sims",
			Plugin:      "workflow-plugin-twilio",
			Description: "List Wireless SIMs.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sims", Type: "array", Description: "List of SIM objects"},
				{Key: "count", Type: "number", Description: "Number of SIMs"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_sim",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a Wireless SIM by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "sim_sid", Label: "SIM SID", Type: schema.FieldTypeString, Description: "SIM SID to fetch", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "SIM SID"},
				{Key: "unique_name", Type: "string", Description: "SIM unique name"},
				{Key: "status", Type: "string", Description: "SIM status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_update_sim",
			Plugin:      "workflow-plugin-twilio",
			Description: "Update a Wireless SIM.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "sim_sid", Label: "SIM SID", Type: schema.FieldTypeString, Description: "SIM SID to update", Required: true},
				{Key: "status", Label: "Status", Type: schema.FieldTypeString, Description: "New SIM status"},
				{Key: "unique_name", Label: "Unique Name", Type: schema.FieldTypeString, Description: "New SIM unique name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "SIM SID"},
				{Key: "status", Type: "string", Description: "Updated SIM status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_fleet",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Wireless Rate Plan (fleet).",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Rate plan SID"},
				{Key: "unique_name", Type: "string", Description: "Rate plan unique name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_send_command",
			Plugin:      "workflow-plugin-twilio",
			Description: "Send a command to a Wireless SIM device.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "command", Label: "Command", Type: schema.FieldTypeString, Description: "Command string to send to the device", Required: true},
				{Key: "sim", Label: "SIM", Type: schema.FieldTypeString, Description: "Target SIM SID or unique name"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Command SID"},
				{Key: "status", Type: "string", Description: "Command status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Pricing ───────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_fetch_pricing",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch voice pricing for a destination number.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "destination_number", Label: "Destination Number", Type: schema.FieldTypeString, Description: "Destination phone number to get pricing for", Required: true},
				{Key: "origination_number", Label: "Origination Number", Type: schema.FieldTypeString, Description: "Origination number for inbound pricing"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "destination_number", Type: "string", Description: "Destination number"},
				{Key: "country", Type: "string", Description: "Country name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_usage_records",
			Plugin:      "workflow-plugin-twilio",
			Description: "List usage records for the account.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "category", Label: "Category", Type: schema.FieldTypeString, Description: "Usage category filter (e.g. calls, sms, recordings)"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "records", Type: "array", Description: "List of usage record objects"},
				{Key: "count", Type: "number", Description: "Number of records"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Accounts ──────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_list_accounts",
			Plugin:      "workflow-plugin-twilio",
			Description: "List sub-accounts under the master account.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "accounts", Type: "array", Description: "List of account objects"},
				{Key: "count", Type: "number", Description: "Number of accounts"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_api_key",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a new Twilio API Key.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Label for the API key"},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "API key SID"},
				{Key: "friendly_name", Type: "string", Description: "Key label"},
				{Key: "secret", Type: "string", Description: "API key secret (returned only at creation time)"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_api_keys",
			Plugin:      "workflow-plugin-twilio",
			Description: "List API Keys in the account.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "keys", Type: "array", Description: "List of API key objects"},
				{Key: "count", Type: "number", Description: "Number of keys"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Content ───────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_content_template",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio Content Template.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Template name", Required: true},
				{Key: "language", Label: "Language", Type: schema.FieldTypeString, Description: "BCP-47 language tag (e.g. en, fr)", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Content template SID"},
				{Key: "friendly_name", Type: "string", Description: "Template name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_content_templates",
			Plugin:      "workflow-plugin-twilio",
			Description: "List Content Templates.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "templates", Type: "array", Description: "List of content template objects"},
				{Key: "count", Type: "number", Description: "Number of templates"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_content_template",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a Content Template by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "sid", Label: "SID", Type: schema.FieldTypeString, Description: "Content template SID to fetch", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Template SID"},
				{Key: "friendly_name", Type: "string", Description: "Template name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── TrustHub ──────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_trust_product",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a TrustHub Trust Product.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "friendly_name", Label: "Friendly Name", Type: schema.FieldTypeString, Description: "Trust product name", Required: true},
				{Key: "email", Label: "Email", Type: schema.FieldTypeString, Description: "Contact email", Required: true},
				{Key: "policy_sid", Label: "Policy SID", Type: schema.FieldTypeString, Description: "TrustHub policy SID", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Trust product SID"},
				{Key: "friendly_name", Type: "string", Description: "Product name"},
				{Key: "status", Type: "string", Description: "Registration status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_trust_products",
			Plugin:      "workflow-plugin-twilio",
			Description: "List TrustHub Trust Products.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "products", Type: "array", Description: "List of trust product objects"},
				{Key: "count", Type: "number", Description: "Number of products"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_fetch_trust_product",
			Plugin:      "workflow-plugin-twilio",
			Description: "Fetch a TrustHub Trust Product by SID.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "sid", Label: "SID", Type: schema.FieldTypeString, Description: "Trust product SID to fetch", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "sid", Type: "string", Description: "Trust product SID"},
				{Key: "friendly_name", Type: "string", Description: "Product name"},
				{Key: "status", Type: "string", Description: "Registration status"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},

		// ── Assistants ────────────────────────────────────────────────────────
		{
			Type:        "step.twilio_create_assistant",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Twilio AI Assistant.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
				{Key: "name", Label: "Name", Type: schema.FieldTypeString, Description: "Assistant name", Required: true},
			},
			Outputs: []schema.StepOutputDef{
				{Key: "id", Type: "string", Description: "Assistant ID"},
				{Key: "name", Type: "string", Description: "Assistant name"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_list_assistants",
			Plugin:      "workflow-plugin-twilio",
			Description: "List Twilio AI Assistants.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "assistants", Type: "array", Description: "List of assistant objects"},
				{Key: "count", Type: "number", Description: "Number of assistants"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
		{
			Type:        "step.twilio_create_knowledge_base",
			Plugin:      "workflow-plugin-twilio",
			Description: "Create a Knowledge Base for a Twilio AI Assistant.",
			ConfigFields: []schema.ConfigFieldDef{
				moduleField,
			},
			Outputs: []schema.StepOutputDef{
				{Key: "id", Type: "string", Description: "Knowledge base ID"},
				{Key: "error", Type: "string", Description: "Error message if the step failed"},
			},
		},
	}
}
