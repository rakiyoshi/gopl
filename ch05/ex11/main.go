package main

import (
	"fmt"
	"os"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":      {"discrete math"},
	"databases":            {"data structures"},
	"discrete math":        {"intro to programing"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programing languages": {"data structures", "computer organization"},
	"linear algebra":       {"calculus"}, // WARN: circular reference
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string, visited []string)

	visitAll = func(items []string, visited []string) {
		for _, item := range items {
			if contains(visited, item) {
				fmt.Fprintf(os.Stderr, "circular reference detected")
				fmt.Println(visited, item)
				os.Exit(1)
			}
			if !seen[item] {
				seen[item] = true
				visitAll(m[item], append(visited, item))
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys, []string{})
	return order
}

func contains(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}
