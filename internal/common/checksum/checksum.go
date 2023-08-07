package checksum

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// FromFile calculates the SHA256 checksum of a file
func FromFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a new SHA-256 hash
	hash := sha256.New()

	// Copy file contents to the hash
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	// Calculate the checksum value
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
