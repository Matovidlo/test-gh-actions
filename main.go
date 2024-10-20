package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	logger := log.Default()
	logger.SetFlags(0)
	logger.Printf("Hello world: %s\n", os.Getenv("TEST_VAR"))
	f, err := os.Open("simple.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file could not be opened: %v\n", err)
		return
	}

	textBytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file could not be opened: %v\n", err)
		return
	}

	logger.Printf("%s", string(textBytes))
}
