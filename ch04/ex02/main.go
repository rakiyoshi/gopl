package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	algorithmStr := *flag.String("argorythm", "SHA256", "SHA256, SHA384, SHA512")

	for {
		var input string
		fmt.Scanf("%s", &input)
		if algorithmStr == "SHA256" {
			fmt.Printf("%s: %x\n", input, sha256.Sum256([]byte(input)))
		} else if algorithmStr == "SHA384" {
			fmt.Printf("%s: %x\n", input, sha512.Sum384([]byte(input)))
		} else if algorithmStr == "SHA512" {
			fmt.Printf("%s: %x\n", input, sha512.Sum512([]byte(input)))
		}
	}
}
