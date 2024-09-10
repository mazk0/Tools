package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
)

func main() {
	mode := flag.String("mode", "decode", "Mode: encode or decode")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		log.Fatalf("Usage: provide exactly one argument, got %d.", len(args))
	}

	switch *mode {
	case "encode":
		encodeBase64(args[0])
	case "decode":
		decodeBase64(args[0])
	default:
		log.Fatalf("Invalid mode: %s. Use 'encode' or 'decode'.", *mode)
	}
}

func decodeBase64(b64 string) {
	data, err := base64.StdEncoding.DecodeString(b64)

	if err != nil {
		log.Fatalf("Error decoding base64: %v\n", err)
	}

	fmt.Println(string(data))
}

func encodeBase64(str string) {
	data := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println(data)
}
