package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type TestLogger struct{}

func NewTestLogger(fileName string) *Logger {
	setOutputFile(fileName)
	return &Logger{}
}

func setOutputFile(fileName string) {

	logFile, _ := os.OpenFile("test_actual_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiLogFile)
	log.SetFlags(log.Flags() &^ log.LstdFlags)

}
func (*TestLogger) PrintStartLog(rootPath string) {
	fmt.Println()
	message := generateMessage("start walking path: %s and delete .DS_Store .....", rootPath)
	printLog(message)
}

func (*TestLogger) PrintDirLog(filePath string) {
	message := generateMessage("path: [%s] (dir).....skip", filePath)
	printLog(message)
}
func (*TestLogger) PrintNormFileLog(filePath string) {
	message := generateMessage("path: [%s] .....skip", filePath)
	printLog(message)
}
func (*TestLogger) PrintTargetFileLog(filePath string) {
	message := generateMessage("path: [%s] .....delete", filePath)
	printLog(message)
}
