package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panic if x == 0
	defer fmt.Printf("defer %d\n", x)
	// nolint:staticcheck
	f(x - 1)
}
