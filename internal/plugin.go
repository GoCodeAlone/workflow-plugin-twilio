// Package internal implements the workflow-plugin-twilio plugin.
package internal

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// twilioPlugin implements sdk.PluginProvider, sdk.ModuleProvider, and sdk.StepProvider.
type twilioPlugin struct{}

// NewTwilioPlugin returns a new twilioPlugin instance.
func NewTwilioPlugin() sdk.PluginProvider {
	return &twilioPlugin{}
}

// Manifest returns plugin metadata.
func (p *twilioPlugin) Manifest() sdk.PluginManifest {
	return sdk.PluginManifest{
		Name:        "workflow-plugin-twilio",
		Version:     "0.1.0",
		Author:      "GoCodeAlone",
		Description: "Twilio communications platform plugin (~90 step types across all Twilio APIs)",
	}
}

// ModuleTypes returns the module type names this plugin provides.
func (p *twilioPlugin) ModuleTypes() []string {
	return []string{"twilio.provider"}
}

// CreateModule creates a module instance of the given type.
func (p *twilioPlugin) CreateModule(typeName, name string, config map[string]any) (sdk.ModuleInstance, error) {
	switch typeName {
	case "twilio.provider":
		m, err := newTwilioModule(name, config)
		if err != nil {
			return nil, err
		}
		return m, nil
	default:
		return nil, fmt.Errorf("twilio plugin: unknown module type %q", typeName)
	}
}

// StepTypes returns the step type names this plugin provides.
func (p *twilioPlugin) StepTypes() []string {
	return allStepTypes()
}

// CreateStep creates a step instance of the given type.
func (p *twilioPlugin) CreateStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	return createStep(typeName, name, config)
}
