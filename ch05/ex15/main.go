package main

import (
	"fmt"
	"log"
	"math"
)

func max(vals ...int) int {
	if len(vals) == 0 {
		log.Fatal("argument must be specified")
	}
	maxValue := math.MinInt
	for _, val := range vals {
		if maxValue < val {
			maxValue = val
		}
	}
	return maxValue
}

func min(vals ...int) int {
	if len(vals) == 0 {
		log.Fatal("argument must be specified")
	}
	minValue := math.MaxInt
	for _, val := range vals {
		if minValue > val {
			minValue = val
		}
	}
	return minValue
}

func main() {
	fmt.Println(max(3))
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(min(3))
	fmt.Println(min(1, 2, 3, 4))
	fmt.Println(max())
}
