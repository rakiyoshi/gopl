package main

import (
	"bytes"
	"unicode"
)

func compressSpace(b []byte) []byte {
	var buf bytes.Buffer
	r := string(b)
	wasSpace := false
	for _, v := range r {
		if unicode.IsSpace(v) {
			if !wasSpace {
				buf.WriteByte(' ')
			}
			wasSpace = true
		} else {
			buf.WriteRune(v)
			wasSpace = false
		}
	}
	return buf.Bytes()
}
