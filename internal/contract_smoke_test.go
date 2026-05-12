package internal

import (
	"reflect"
	"testing"
)

// TestNewTwilioPlugin_ReturnsNonNil is a contract smoke test that verifies the
// exported factory function returns a non-nil PluginProvider. This guards
// against accidental nil-return regressions in the factory and is required by
// the strict-contracts gate in CI.
func TestNewTwilioPlugin_ReturnsNonNil(t *testing.T) {
	p := NewTwilioPlugin()
	if p == nil {
		t.Fatal("NewTwilioPlugin() returned nil")
	}
	// Guard against typed-nil (interface non-nil but underlying pointer is nil),
	// which would panic at sdk.Serve call time.
	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		t.Fatal("NewTwilioPlugin() returned a typed-nil interface value")
	}
	// Type-assert to concrete type to confirm factory wiring.
	if _, ok := p.(*twilioPlugin); !ok {
		t.Fatalf("NewTwilioPlugin() returned unexpected type %T", p)
	}
}
