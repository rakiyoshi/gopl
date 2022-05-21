package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	mid := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Now()
	fmt.Printf("Ineffective: %v\n", mid.Sub(start).String())
	fmt.Printf("Effective  : %v\n", end.Sub(mid).String())
}
