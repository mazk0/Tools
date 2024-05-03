package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	switch {
	case len(args) < 1:
		fmt.Println("Usage: provide a base64 encoded string as the only argument.")
	case len(args) > 1:
		fmt.Printf("Error: Only one argument expected, got %d.\n", len(args))
	default:
		decodeBase64(args[0])
	}
}

func decodeBase64(b64 string) {
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		fmt.Printf("Error decoding base64: %v\n", err)
		return
	}
	fmt.Println(string(data))
}
