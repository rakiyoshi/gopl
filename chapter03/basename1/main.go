package main

import "fmt"

func main() {
	fmt.Println(basename("a/b/c.go"))
}
func basename(s string) string {
	// remove '/' and before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// hold all before '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
