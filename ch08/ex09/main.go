package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// retrieve file trees
	fileSizes := make(chan struct{ iRoot, size int64 })
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, int64(i), &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles, nbytes := make([]int64, len(roots)), make([]int64, len(roots))
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes has been closed
			}
			nfiles[size.iRoot]++
			nbytes[size.iRoot] += size.size
		case <-tick:
			printDiskUsage(nfiles, nbytes, roots)
		}
	}
	fmt.Println("---")
	printDiskUsage(nfiles, nbytes, roots)
}

func printDiskUsage(nfiles, nbytes []int64, roots []string) {
	for i, root := range roots {
		fmt.Printf("%s: %d files %.1f GB\n", root, nfiles[i], float64(nbytes[i])/1e9)
	}
}

func walkDir(dir string, iRoot int64, n *sync.WaitGroup, fileSizes chan<- struct{ iRoot, size int64 }) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, iRoot, n, fileSizes)
		} else {
			fileSizes <- struct{ iRoot, size int64 }{iRoot, entry.Size()}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex09: %v\n", err)
		return nil
	}
	infos := make([]os.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex09: %v\n", err)
		}
		infos = append(infos, info)
	}
	return infos
}
