package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		err := ex07(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
	}
}

// TODO: fix me and add test
func ex07(url string) error {
	doc, err := getURL(url)
	if err != nil {
		return err
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

func getURL(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		var attrs string
		for _, attr := range n.Attr {
			attrs += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
		}
		fmt.Printf("%*s<%s%s>", depth*2, "", n.Data, attrs)
		if n.FirstChild != nil && n.FirstChild.Type != html.TextNode {
			fmt.Printf("\n")
		}
		depth++
	case html.TextNode:
		fmt.Printf("%s", n.Data)
	case html.CommentNode:
		fmt.Printf("%s", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			if n.FirstChild.Type != html.TextNode {
				fmt.Printf("%*s", depth*2, "")
			}
			fmt.Printf("</%s>\n", n.Data)
		} else {
			fmt.Print("\n")
		}
	}
}
