package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToSting([]int{1, 2, 3}))
}

func intsToSting(values []int) string {
	var buf bytes.Buffer
	err := buf.WriteByte('[')
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
