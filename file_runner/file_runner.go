package file_runner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/yamanoi-shu/go_ds_store/file_runner/repository"
)

type FileRunner struct {
	RootPath string
	logger   repository.LoggerRepository
}

func NewFileRunner(rootPath string, logger repository.LoggerRepository) (*FileRunner, error) {

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
	fr.logger.PrintStartLog(fr.RootPath)

	err = filepath.Walk(fr.RootPath, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if info.IsDir() {
			fr.logger.PrintDirLog(path)
			return nil
		}
		if info.Name() == ".DS_Store" {
			fr.logger.PrintTargetFileLog(path)
			err = os.Remove(path)
			return err
		}
		fr.logger.PrintNormFileLog(path)
		return nil
	})

	return
}
