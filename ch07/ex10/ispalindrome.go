package ex10

import "sort"

func IsPalindrome(s sort.Interface) bool {
	var i, j int
	if s.Len()%2 == 0 {
		i = (s.Len() - 1) / 2
		j = i + 1
	} else {
		i = (s.Len()-1)/2 - 1
		j = i + 2
	}
	for ; i >= 0 && j < s.Len(); i, j = i-1, j+1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
