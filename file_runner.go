package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type FileRunner struct {
	RootPath string
	logger   *Logger
}

func NewFileRunner(rootPath string, logger *Logger) (*FileRunner, error) {

	fInfo, err := os.Stat(rootPath)

	if err != nil {
		err := fmt.Errorf("this file path ( %s ) is invalid !", rootPath)
		return nil, err
	}

	if !fInfo.IsDir() {
		err := fmt.Errorf("this file ( %s ) is not directory !", rootPath)
		return nil, err
	}

	return &FileRunner{
		RootPath: rootPath,
		logger:   logger,
	}, nil
}

func (fr *FileRunner) DeleteAllDSStore() (err error) {
	fr.logger.printStartLog(fr.RootPath)

	err = filepath.Walk(fr.RootPath, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if info.IsDir() {
			fr.logger.printDirLog(path)
			return nil
		}
		if info.Name() == ".DS_Store" {
			fr.logger.printDSStoreLog(path)
			err = os.Remove(path)
			return err
		}
		fr.logger.printNormFileLog(path)
		return nil
	})

	return
}
