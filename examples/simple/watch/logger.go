package main

import (
	"fmt"
)

// LogMetrics logs performance stats key/value pairs.
func LogMetrics(kvp [][]string) {
	for _, kv := range kvp {
		var values string
		for i, val := range kv {
			switch i {
			case 0:
				continue
			case 1:
				values += val
			default:
				values += "("
				values += val
				values += ")"
			}
		}
		fmt.Printf("%s -> %s; ", kv[0], values)
	}
	fmt.Println()
}
