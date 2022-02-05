package main

import (
	"fmt"
	"log"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}
func (*Logger) printStartLog(rootPath string) {
	fmt.Println()
	message := generateMessage("start walking path: %s and delete .DS_Store .....", rootPath)
	printLog(message)
}

func (*Logger) printDirLog(filePath string) {
	message := generateMessage("path: [%s] (dir).....skip", filePath)
	printLog(message)
}
func (*Logger) printNormFileLog(filePath string) {
	message := generateMessage("path: [%s] .....skip", filePath)
	printLog(message)
}
func (*Logger) printDSStoreLog(filePath string) {
	message := generateMessage("path: [%s] .....delete", filePath)
	printLog(message)
}

func generateMessage(format string, args ...string) string {
	return fmt.Sprintf(format, args)
}
func printLog(message string) {
	log.Println(message)
	fmt.Println()
}
