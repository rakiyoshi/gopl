package main

import (
	"bufio"
	"fmt"
	"os"
)

type WordCounts map[string]int

func wordfreq(filename string) WordCounts {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	wc := WordCounts{}
	for input.Scan() {
		word := input.Text()
		wc[word] += 1
	}
	return wc
}
