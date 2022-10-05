package main

import (
	"flag"
	"fmt"
	"gopl/ch05/links"
	"log"
)

var depthFlag int

var tokens = make(chan struct{}, 20)

func init() {
	flag.IntVar(&depthFlag, "depth", -1, "crawl depth")
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()

	type WorkList struct {
		depth int
		list  []string
	}
	worklist := make(chan WorkList)

	type UnseenLinks struct {
		depth int
		list  string
	}
	unseenLinks := make(chan UnseenLinks)

	go func() { worklist <- WorkList{0, flag.Args()} }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.list)
				go func(depth int) { worklist <- WorkList{depth, foundLinks} }(link.depth)
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		if depthFlag >= 0 && list.depth > depthFlag {
			break
		}
		for _, link := range list.list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- UnseenLinks{list.depth + 1, link}
			}
		}
	}
}
