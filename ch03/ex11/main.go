package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))

	fmt.Println(comma("-1234"))
	fmt.Println(comma("-12345"))

	fmt.Println(comma("1234.5678"))
	fmt.Println(comma("-1234.567890123"))
}

func comma(s string) string {
	var buf bytes.Buffer
	var n int
	if len(s) == 0 {
		return ""
	}

	if string(s[0]) == "-" {
		buf.WriteByte('-')
		s = s[1:]
	}

	var s1, s2 string
	if dot := strings.Index(s, "."); dot > 0 {
		s1 = s[:dot]
		s2 = s[dot+1:]
	} else {
		s1 = s
	}

	// before decimal point
	n = len(s1) % 3
	if n == 0 {
		buf.WriteString(s1[:3])
		n += 3
	} else {
		buf.WriteString(s1[:n])
	}

	for n < len(s1) {
		fmt.Fprintf(&buf, ",%s", s[n:n+3])
		n += 3
	}

	// after decimal point
	n = 0
	if len(s2) > 0 {
		buf.WriteByte('.')
	}
	for n+3 < len(s2) {
		fmt.Fprintf(&buf, "%s,", s2[n:n+3])
		n += 3
	}
	buf.WriteString(s2[n:])
	return buf.String()
}
