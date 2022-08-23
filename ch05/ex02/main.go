package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
		os.Exit(1)
	}
	for element, count := range visit(nil, doc) {
		fmt.Printf("% 8s: %d\n", element, count)
	}
}

func visit(elementCount map[string]int64, n *html.Node) map[string]int64 {
	if n.Type == html.ElementNode {
		if elementCount == nil {
			elementCount = map[string]int64{}
		}
		elementCount[n.Data] += 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementCount = visit(elementCount, c)
	}
	return elementCount
}
