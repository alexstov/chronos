package watch

import (
	"fmt"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// LogMetrics logs performance stats key/value pairs.
//
//nolint:forbidigo
func LogMetrics(kvp [][]string) {
	var out string
	for _, kv := range kvp {
		out = fmt.Sprintf("%s -> %s; ", kv[0], kv[1])
	}
	fmt.Println(out)
}
