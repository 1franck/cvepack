package common

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func DirectoryExists(path string) bool {
	if fileInfo, err := os.Stat(path); err == nil && fileInfo.IsDir() {
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

func CopyFile(sourcePath, destinationPath string) error {
	// Open the source file
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create or overwrite the destination file
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copy the contents from the source to the destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func ReadAllFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
