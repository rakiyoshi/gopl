package main

import "unicode/utf8"

func reverse(b []byte) {
	var prev int
	for i := 0; i < utf8.RuneCount(b)-1; i++ {
		_, size := utf8.DecodeRune(b)
		rotateLeft(b[:len(b)-prev], size)
		prev += size
	}
}

func rotateLeft(s []byte, n int) {
	revBytes(s[:n])
	revBytes(s[n:])
	revBytes(s)
}

func revBytes(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
