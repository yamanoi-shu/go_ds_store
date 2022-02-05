package repository

type LoggerRepository interface {
	PrintStartLog(rootPath string)
	PrintDirLog(filePath string)
	PrintNormFileLog(filePath string)
	PrintTargetFileLog(filePath string)
}
