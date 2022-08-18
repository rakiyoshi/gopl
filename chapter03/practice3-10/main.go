package main

import (
	"bytes"
	"fmt"
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
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	if n == 0 {
		buf.WriteString(s[:3])
		n += 3
	} else {
		buf.WriteString(s[:n])
	}
	for n < len(s) {
		fmt.Fprintf(&buf, ",%s", s[n:n+3])
		n += 3
	}
	return buf.String()
}
