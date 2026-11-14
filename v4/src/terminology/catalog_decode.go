package terminology

import (
	"encoding/json"
	"fmt"
)

// mustCatalog decodes an embedded catalog literal. Panics on invalid data: a broken
// catalog must fail at start-up, not silently return an empty map at runtime.
func mustCatalog(raw string) map[string]any {
	var hasil map[string]any
	if err := json.Unmarshal([]byte(raw), &hasil); err != nil {
		panic(fmt.Sprintf("catalog tidak sah: %v", err))
	}
	return hasil
}
