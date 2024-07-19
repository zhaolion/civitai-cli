package util

import (
	"fmt"
	"os"
)

func CheckTargetDirExists(targetDir string) error {
	stat, err := os.Stat(targetDir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("target directory does not exist: %s", targetDir)
		} else {
			return err
		}
	}

	if !stat.IsDir() {
		return fmt.Errorf("target directory is not a directory: %s", targetDir)
	}

	return nil
}

func CheckTargetFileNotExists(targetFile string) error {
	_, err := os.Stat(targetFile)
	if err == nil {
		return fmt.Errorf("target file exists: %s", targetFile)
	}
	if !os.IsNotExist(err) {
		return err
	}

	return nil
}
