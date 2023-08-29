package update

import (
	"archive/zip"
	"bufio"
	"cvepack/core"
	"cvepack/core/common"
	"cvepack/core/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	ErrorDatabaseFolderNotFound            core.ErrorMsg = "database folder not found"
	ErrorDatabaseFileNotFound              core.ErrorMsg = "database file not found"
	ErrorDatabaseChecksumFileNotFound      core.ErrorMsg = "database checksum file not found"
	ErrorDatabaseServerChecksumFileInvalid core.ErrorMsg = "error checking server database checksum: %s"
	ErrorDatabaseReadingLocalChecksum      core.ErrorMsg = "error reading local database checksum: %s"
	ErrorDatabaseChecksumMismatch          core.ErrorMsg = "databases checksums mismatch"
)

func UpdateDatabase(outputDir string) {

	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	// Download the ZIP file
	resp, err := http.Get(config.Default.DatabaseUrl)
	if err != nil {
		fmt.Println("Error downloading ZIP file:", err)
		return
	}
	defer resp.Body.Close()

	// Create a temporary file to store the downloaded content
	tmpFile, err := os.CreateTemp("", "downloaded-*.zip")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Copy the downloaded content to the temporary file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		fmt.Println("Error copying content to temporary file:", err)
		return
	}

	// Open the downloaded ZIP file
	zipReader, err := zip.OpenReader(tmpFile.Name())
	if err != nil {
		fmt.Println("Error opening ZIP file:", err)
		return
	}
	defer zipReader.Close()

	// Extract the contents
	for _, file := range zipReader.File {
		filePath := filepath.Join(outputDir, file.Name)

		if file.FileInfo().IsDir() {
			// Create directories
			err := os.MkdirAll(filePath, file.Mode())
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		} else {
			// Create files
			outFile, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer outFile.Close()

			rc, err := file.Open()
			if err != nil {
				fmt.Println("Error opening file from ZIP:", err)
				return
			}
			defer rc.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				fmt.Println("Error extracting file:", err)
				return
			}
		}
	}
}

func IsAvailable(config config.Config) (bool, core.ErrorMsg) {
	if !common.DirectoryExists(config.DatabaseRootDir) {
		return true, ErrorDatabaseFolderNotFound
	} else if !common.FileExists(config.DatabaseFilePath()) {
		return true, ErrorDatabaseFileNotFound
	} else if !common.FileExists(config.DatabaseChecksumFilePath()) {
		return true, ErrorDatabaseChecksumFileNotFound
	}

	resp, err := http.Get(config.DatabaseChecksumUrl)
	if err != nil {
		return false, ErrorDatabaseServerChecksumFileInvalid.Sprintf(err)
	}
	defer resp.Body.Close()

	dbChecksum := ""
	scanner := bufio.NewScanner(resp.Body)
	if scanner.Scan() {
		dbChecksum = scanner.Text()
	}

	localChecksum := ""
	localChecksumBytes, err := common.ReadAllFile(config.DatabaseChecksumFilePath())
	if err != nil {
		return false, ErrorDatabaseReadingLocalChecksum.Sprintf(err)
	}
	localChecksum = string(localChecksumBytes)

	if dbChecksum != localChecksum {
		return true, ErrorDatabaseChecksumMismatch
	}

	return false, core.EmptyError
}
