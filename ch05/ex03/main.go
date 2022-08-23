package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex03: %v\n", err)
		os.Exit(1)
	}
	for _, t := range visit(nil, doc) {
		fmt.Printf("'%s'\n", t)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style" || n.Data == "iframe") {
		return texts
	}
	if n.Type == html.TextNode {
		replaced := strings.ReplaceAll(n.Data, " ", "")
		replaced = strings.ReplaceAll(replaced, "\t", "")
		replaced = strings.ReplaceAll(replaced, "\n", "")
		if len(replaced) > 0 {
			texts = append(texts, trimPrefixSuffixAll(n.Data))
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = visit(texts, c)
	}
	return texts
}

func trimPrefixSuffixAll(s string) string {
	for {
		if c := s[0:1]; c == " " || c == "\n" {
			s = strings.TrimPrefix(s, c)
		} else {
			break
		}
	}
	for {
		if c := s[len(s)-1:]; c == " " || c == "\n" {
			s = strings.TrimSuffix(s, c)
		} else {
			break
		}
	}
	return s
}
