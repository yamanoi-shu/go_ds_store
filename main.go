package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <root path>")
		return
	}

	filePath := args[0]

	logger := NewLogger()
	fr, err := NewFileRunner(filePath, logger)
	if err != nil {
		log.Fatal(err)
	}

	fr.DeleteAllDSStore()
}
