package main

import "strings"

func expand(s string, f func(string) string) string {
	var results []string
	for _, word := range strings.Fields(s) {
		if string(word[0]) == "$" {
			results = append(results, f(word[1:]))
		} else {
			results = append(results, word)
		}
	}
	return strings.Join(results, " ")
}
