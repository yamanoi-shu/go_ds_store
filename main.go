package main

import (
	"flag"
	"fmt"
	"go_ds_store/file_runner"
	"go_ds_store/file_runner/repository"
	"go_ds_store/logger"
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

	var logHandler repository.LoggerRepository

	if os.Getenv("LOG_FILE") == "true" {
		logHandler = logger.NewTestLogger("test_actual_log.txt")
	} else {
		logHandler = logger.NewLogger()
	}

	filePath := args[0]

	fr, err := file_runner.NewFileRunner(filePath, logHandler)
	if err != nil {
		log.Fatal(err)
		return
	}

	fr.DeleteAllDSStore()
}
