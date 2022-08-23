package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex05: %v\n", err)
	}
	fmt.Printf("words: %d\n", words)
	fmt.Printf("images: %d\n", images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		replaced := trimPrefixSuffixAll(n.Data)
		words += len(strings.Split(replaced, " "))
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images += 1
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func trimPrefixSuffixAll(s string) string {
	for {
		if c := s[0:1]; c == " " || c == "\n" {
			s = strings.TrimPrefix(s, c)
			if len(s) == 0 {
				return ""
			}
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
