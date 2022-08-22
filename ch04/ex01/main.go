package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c1: %08b\n", c1)
	fmt.Printf("c2: %08b\n", c2)
	var diff int
	for i := 0; i < len(c1); i++ {
		diff += int(pc[c1[i]^c2[i]])
	}
	fmt.Printf("diff: %d", diff)
}
