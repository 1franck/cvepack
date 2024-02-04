package database

import (
	"cvepack/core/sqlite"
	"errors"
	"os"
)

func IsDatabaseOk(dbPath string) error {
	db, err := sqlite.Connect(dbPath)
	defer db.Close()
	if err != nil {
		return err
	}

	vulCount, err := CountVulnerabilities(db)
	if err != nil {
		return err
	}
	if vulCount == 0 {
		return errors.New("no vulnerabilities found in database :/")
	}
	return nil
}

func LastModified(dbPath string) (string, error) {
	fileInfo, err := os.Stat(dbPath)
	if err != nil {
		return "", err
	}
	return fileInfo.ModTime().Format("2006-01-02 15:04:05"), nil
}
