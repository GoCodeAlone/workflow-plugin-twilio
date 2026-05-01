// gen-schemas generates the stepSchemas JSON block for plugin.json.
// Usage: go run ./tools/gen-schemas > /tmp/step_schemas.json
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
