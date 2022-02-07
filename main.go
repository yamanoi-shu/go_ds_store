package main

import (
	"flag"
	"fmt"
	"go_ds_store/file_runner"
	"go_ds_store/logger"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go <root path>")
		return
	}

	if os.Getenv("LOG_FILE") == "true" {
		logFile, _ := os.Open("test.txt")
		multiLog := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiLog)
	}

	filePath := args[0]

	logger := logger.NewLogger()
	fr, err := file_runner.NewFileRunner(filePath, logger)
	if err != nil {
		log.Fatal(err)
		return
	}

	fr.DeleteAllDSStore()
}
