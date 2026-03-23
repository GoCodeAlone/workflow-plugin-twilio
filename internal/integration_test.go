package internal_test

import (
	"testing"

	"github.com/GoCodeAlone/workflow/wftest"
)

func TestIntegration_SendSMS(t *testing.T) {
	h := wftest.New(t, wftest.WithYAML(`
pipelines:
  notify:
    steps:
      - name: send
        type: step.twilio_send_sms
        config:
          to: "+15555555555"
          body: "hello"
      - name: confirm
        type: step.set
        config:
          values:
            sent: true
`),
		wftest.MockStep("step.twilio_send_sms", wftest.Returns(map[string]any{
			"sid": "SM123", "status": "queued",
		})),
	)
	result := h.ExecutePipeline("notify", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["sent"] != true {
		t.Error("expected sent=true")
	}
	if result.Output["sid"] != "SM123" {
		t.Errorf("expected sid=SM123, got %v", result.Output["sid"])
	}
}

func TestIntegration_SendVerification(t *testing.T) {
	h := wftest.New(t, wftest.WithYAML(`
pipelines:
  verify:
    steps:
      - name: send_otp
        type: step.twilio_send_verification
        config:
          service_sid: "VAabc123"
          to: "+15555555555"
          channel: "sms"
      - name: mark_sent
        type: step.set
        config:
          values:
            otp_sent: true
`),
		wftest.MockStep("step.twilio_send_verification", wftest.Returns(map[string]any{
			"sid":    "VEabc123",
			"status": "pending",
			"to":     "+15555555555",
		})),
	)
	result := h.ExecutePipeline("verify", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["otp_sent"] != true {
		t.Error("expected otp_sent=true")
	}
	if result.Output["status"] != "pending" {
		t.Errorf("expected status=pending, got %v", result.Output["status"])
	}
}

func TestIntegration_CreateCall(t *testing.T) {
	rec := wftest.RecordStep("step.twilio_create_call")
	rec.WithOutput(map[string]any{
		"sid":    "CAabc123",
		"status": "queued",
		"to":     "+15555555555",
		"from":   "+18005551234",
	})
	h := wftest.New(t, wftest.WithYAML(`
pipelines:
  call:
    steps:
      - name: dial
        type: step.twilio_create_call
        config:
          to: "+15555555555"
          from: "+18005551234"
          url: "https://example.com/twiml"
      - name: done
        type: step.set
        config:
          values:
            call_placed: true
`),
		rec,
	)
	result := h.ExecutePipeline("call", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["call_placed"] != true {
		t.Error("expected call_placed=true")
	}
	if rec.CallCount() != 1 {
		t.Errorf("expected step called once, got %d", rec.CallCount())
	}
	calls := rec.Calls()
	if calls[0].Config["to"] != "+15555555555" {
		t.Errorf("expected to=+15555555555, got %v", calls[0].Config["to"])
	}
}
