package main

import (
	"flag"
	"fmt"
	tempconv "gopl/ch07/tempconv_ex06"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
