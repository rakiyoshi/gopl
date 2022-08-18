package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(isAnagram("live", "evil"))
	fmt.Println(isAnagram("fuck", "you"))
	fmt.Println(isAnagram("animal", "inimal"))
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Strings, s2Strings := strings.Split(s1, ""), strings.Split(s2, "")
	sort.Strings(s1Strings)
	sort.Strings(s2Strings)
	for i := 0; i < len(s1Strings); i++ {
		if s1Strings[i] != s2Strings[i] {
			return false
		}
	}
	return true
}
