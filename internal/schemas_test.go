package internal

import (
	"testing"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// TestSchemaCoverage verifies that every advertised step type has a corresponding
// strict contract descriptor (StepSchema), and that the module schema descriptor
// is present for every advertised module type.
func TestSchemaCoverage(t *testing.T) {
	p := &twilioPlugin{}

	// ── Step coverage ─────────────────────────────────────────────────────────
	stepTypes := p.StepTypes()
	schemas := allStepSchemas()

	schemasByType := make(map[string]bool, len(schemas))
	for _, s := range schemas {
		if s.Type == "" {
			t.Error("found StepSchema with empty type")
		}
		schemasByType[s.Type] = true
	}

	for _, st := range stepTypes {
		if !schemasByType[st] {
			t.Errorf("step type %q has no StepSchema contract descriptor", st)
		}
	}
	for schemaType := range schemasByType {
		found := false
		for _, st := range stepTypes {
			if st == schemaType {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("StepSchema %q has no matching step type (orphaned descriptor)", schemaType)
		}
	}

	t.Logf("step coverage: %d/%d have descriptors", len(schemas), len(stepTypes))

	// ── Module coverage ───────────────────────────────────────────────────────
	// Verify SchemaProvider is implemented and each module type has a descriptor.
	sp, ok := (interface{})(p).(sdk.SchemaProvider)
	if !ok {
		t.Fatal("twilioPlugin does not implement sdk.SchemaProvider")
	}
	moduleSchemas := sp.ModuleSchemas()
	moduleSchemasByType := make(map[string]bool, len(moduleSchemas))
	for _, ms := range moduleSchemas {
		if ms.Type == "" {
			t.Error("found ModuleSchemaData with empty type")
		}
		moduleSchemasByType[ms.Type] = true
	}

	for _, mt := range p.ModuleTypes() {
		if !moduleSchemasByType[mt] {
			t.Errorf("module type %q has no ModuleSchemaData contract descriptor", mt)
		}
	}

	t.Logf("module coverage: %d/%d have descriptors", len(moduleSchemas), len(p.ModuleTypes()))
}

// TestModuleSchemaFields verifies that every ModuleSchemaData has the expected
// fields populated (type, label, category, description, at least one config field).
func TestModuleSchemaFields(t *testing.T) {
	p := &twilioPlugin{}
	sp, ok := (interface{})(p).(sdk.SchemaProvider)
	if !ok {
		t.Fatal("twilioPlugin does not implement sdk.SchemaProvider")
	}
	for _, ms := range sp.ModuleSchemas() {
		if ms.Type == "" {
			t.Errorf("ModuleSchemaData has empty Type")
		}
		if ms.Label == "" {
			t.Errorf("ModuleSchemaData %q has empty Label", ms.Type)
		}
		if ms.Category == "" {
			t.Errorf("ModuleSchemaData %q has empty Category", ms.Type)
		}
		if ms.Description == "" {
			t.Errorf("ModuleSchemaData %q has empty Description", ms.Type)
		}
		if len(ms.ConfigFields) == 0 {
			t.Errorf("ModuleSchemaData %q has no ConfigFields", ms.Type)
		}
	}
}

// TestStepSchemaFields verifies that every StepSchema has required fields and
// at least one output defined.
func TestStepSchemaFields(t *testing.T) {
	for _, s := range allStepSchemas() {
		if s.Type == "" {
			t.Error("StepSchema has empty Type")
		}
		if s.Description == "" {
			t.Errorf("StepSchema %q has empty Description", s.Type)
		}
		if s.Plugin == "" {
			t.Errorf("StepSchema %q has empty Plugin", s.Type)
		}
		if len(s.Outputs) == 0 {
			t.Errorf("StepSchema %q has no Outputs defined", s.Type)
		}
	}
}
