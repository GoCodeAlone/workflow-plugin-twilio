package internal

import "testing"

// TestNewTwilioPlugin_ReturnsNonNil is a contract smoke test that verifies the
// exported factory function returns a non-nil PluginProvider. This guards
// against accidental nil-return regressions in the factory and is required by
// the strict-contracts gate in CI.
func TestNewTwilioPlugin_ReturnsNonNil(t *testing.T) {
	p := NewTwilioPlugin()
	if p == nil {
		t.Fatal("NewTwilioPlugin() returned nil")
	}
}
