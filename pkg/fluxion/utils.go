package fluxion

import (
	"fmt"

	"encoding/json"
)

type allocation struct {
	Type      string
	Name      string
	Basename  string
	CoreCount int
}

// parseAllocResult takes an allocated (string) and parses into a list of allocation
// We include the pod namespace/name for debugging later
// TODO this needs refactor for non Kubernetes case
func parseAllocResult(allocated string) []allocation {
	var dat map[string]interface{}
	result := []allocation{}

	// Keep track of total core count across allocated
	corecount := 0

	// This should not happen - the string we get back should parse.
	if err := json.Unmarshal([]byte(allocated), &dat); err != nil {
		panic(err)
	}
	// Parse graph and nodes into interfaces
	// TODO look at github.com/mitchellh/mapstructure
	// that might make this easier
	nodes := dat["graph"]
	str1 := nodes.(map[string]interface{})
	str2 := str1["nodes"].([]interface{})

	for _, item := range str2 {
		str1 = item.(map[string]interface{})
		metadata := str1["metadata"].(map[string]interface{})
		if metadata["type"].(string) == "core" {
			corecount = corecount + 1
		}
		if metadata["type"].(string) == "node" {
			result = append(result, allocation{
				Type:      metadata["type"].(string),
				Name:      metadata["name"].(string),
				Basename:  metadata["basename"].(string),
				CoreCount: corecount,
			})

			// Reset the corecount once we've added to a node
			corecount = 0
		}
	}
	fmt.Println("Final node result")
	for i, alloc := range result {
		fmt.Printf("Node %d: %s\n", i, alloc.Name)
		fmt.Printf("  Type: %s\n  Name: %s\n  Basename: %s\n  CoreCount: %d\n",
			alloc.Type, alloc.Name, alloc.Basename, alloc.CoreCount)

	}
	return result
}

// Utility functions
func printOutput(reserved bool, allocated string, at int64, overhead float64, jobid uint64, fluxerr error) {
	fmt.Println("\n\t----Match Allocate output---")
	fmt.Printf("jobid: %d\nreserved: %t\nallocated: %s\nat: %d\noverhead: %f\n", jobid, reserved, allocated, at, overhead)

	// Only print error if we had one
	if fluxerr != nil {
		fmt.Printf("error: %w\n", fluxerr)
	}
}
