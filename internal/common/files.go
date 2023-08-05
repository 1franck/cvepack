package common

import (
	"errors"
	"fmt"
	"os"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func ValidateDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("Path %s does not exist", path))
	}

	if fileInfo, err := os.Stat(path); err == nil && !fileInfo.IsDir() {
		return errors.New(fmt.Sprintf("Path %s is not a directory", path))
	}

	return nil
}
