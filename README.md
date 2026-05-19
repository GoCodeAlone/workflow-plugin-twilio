# workflow-plugin-twilio

> ✅ **Verified** — used in production at **buymywishlist**. This plugin has been validated end-to-end in a merged main-branch wfctl.yaml of an active GoCodeAlone project.

Twilio communications platform plugin for the [GoCodeAlone/workflow](https://github.com/GoCodeAlone/workflow) engine. Provides ~90 step types across SMS, MMS, WhatsApp, voice, video, verify, conversations, and more.

## Module Types

- `twilio.provider` — Twilio REST client (account SID + auth token or API key)

## Key Step Types

| Step | Description |
|------|-------------|
| `step.twilio_send_sms` | Send an SMS message |
| `step.twilio_send_mms` | Send an MMS message |
| `step.twilio_send_whatsapp` | Send a WhatsApp message |
| `step.twilio_send_verification` | Start a Verify OTP flow |
| `step.twilio_check_verification` | Verify an OTP code |
| `step.twilio_create_call` | Initiate a voice call |
| `step.twilio_create_conversation` | Create a Conversations conversation |

See `plugin.json` for the full list of ~90 step types.

## Quick start

See [`examples/minimal/config.yaml`](examples/minimal/config.yaml) for a minimal working configuration.

## Requirements

- workflow engine ≥ `0.51.2`
- `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` environment variables

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).
