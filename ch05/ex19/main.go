package main

import (
	"fmt"
)

func main() {
	fmt.Println(noReturn(1))
	fmt.Println(noReturn(2))
}

// nolint:typecheck
func noReturn(x int) (val int) {
	defer func() {
		p := recover()
		if fmt.Sprintf("%T", p) == "int" {
			val = x
		}

	}()
	panic(x)
}
