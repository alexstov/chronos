package watch

import "fmt"

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// LogMetrics logs performance stats key/value pairs.
func LogMetrics(kvp [][]string) {
	for _, kv := range kvp {
		fmt.Printf("%s -> %s; ", kv[0], kv[1])
	}
	fmt.Println()
}
