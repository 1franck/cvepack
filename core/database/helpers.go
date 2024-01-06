package database

import (
	"cvepack/core/sqlite"
	"errors"
	"os"
	"time"
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
	modTime := fileInfo.ModTime()
	return modTime.Format(time.RFC1123), nil
}
