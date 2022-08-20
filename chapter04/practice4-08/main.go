package main

import (
	"unicode"
	"unicode/utf8"
)

type Counts struct {
	control     int
	letter      int
	mark        int
	number      int
	punctuation int
	space       int
	symbol      int
}

func charcount(str string) Counts {
	b := []byte(str)
	var counts Counts
	var pos int
	for pos < len(b) {
		r, size := utf8.DecodeRune(b[pos:])

		if unicode.IsControl(r) {
			counts.control += 1
		}
		if unicode.IsLetter(r) {
			counts.letter += 1
		}
		if unicode.IsMark(r) {
			counts.mark += 1
		}
		if unicode.IsNumber(r) {
			counts.number += 1
		}
		if unicode.IsPunct(r) {
			counts.punctuation += 1
		}
		if unicode.IsSpace(r) {
			counts.space += 1
		}
		if unicode.IsSymbol(r) {
			counts.symbol += 1
		}
		pos += size
	}
	return counts
}
