// Fetch prints contents of the URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		for {
			n, err := io.Copy(os.Stdout, resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
			if n == 0 {
				break
			}
		}
	}
}
