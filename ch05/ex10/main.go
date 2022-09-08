package main

import (
	"fmt"
)

var prereqs = map[string]map[string]struct{}{
	"algorithms": {"data structures": struct{}{}},
	"calculus":   {"linear algebra": struct{}{}},
	"compilers": {
		"data structures":       struct{}{},
		"formal languages":      struct{}{},
		"computer organization": struct{}{},
	},
	"data structures":  {"discrete math": struct{}{}},
	"databases":        {"data structures": struct{}{}},
	"discrete math":    {"intro to programing": struct{}{}},
	"formal languages": {"discrete math": struct{}{}},
	"networks":         {"operating systems": struct{}{}},
	"operating systems": {
		"data structures":       struct{}{},
		"computer organization": struct{}{},
	},
	"programing languages": {
		"data structures":       struct{}{},
		"computer organization": struct{}{},
	},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]struct{}) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]struct{})

	visitAll = func(items map[string]struct{}) {
		for k := range items {
			if !seen[k] {
				seen[k] = true
				visitAll(m[k])
				order = append(order, k)
			}
		}
	}

	keys := make(map[string]struct{})
	for key := range m {
		keys[key] = struct{}{}
	}

	visitAll(keys)
	return order
}
