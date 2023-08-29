package database

import (
	"cvepack/core/sqlite"
	"errors"
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
