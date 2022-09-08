package main

import (
	"fmt"
	"strings"
)

func join(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}

func main() {
	fmt.Println(join(".", "192", "168", "11", "1"))
}
