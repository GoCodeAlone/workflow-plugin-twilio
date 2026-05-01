// gen-schemas regenerates the stepSchemas JSON block for plugin.json.
//
// Usage:
//
//	go run ./tools/gen-schemas > /tmp/step_schemas.json
//
// To update plugin.json, replace the "stepSchemas" value with the output of
// this tool. The expected workflow is:
//
//  1. Edit allStepSchemas() in internal/schemas.go
//  2. Run: go run ./tools/gen-schemas > /tmp/step_schemas.json
//  3. Update the "stepSchemas" key in plugin.json with the generated JSON array
package main

import (
"encoding/json"
"os"
"sort"

"github.com/GoCodeAlone/workflow-plugin-twilio/internal"
)

func main() {
schemas := internal.AllStepSchemas()
sort.Slice(schemas, func(i, j int) bool { return schemas[i].Type < schemas[j].Type })
enc := json.NewEncoder(os.Stdout)
enc.SetIndent("", "    ")
if err := enc.Encode(schemas); err != nil {
panic(err)
}
}
