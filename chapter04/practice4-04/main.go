package main

func rotate(s *[]int, n int) {
	*s = append(*s, (*s)[:n]...)[n : len(*s)+n]
}
