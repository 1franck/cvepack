package search

import (
	"database/sql"
	"fmt"
	"log"
)

func LookupPackageQuerier(db *sql.DB) *lookupPackageQuerier {
	return &lookupPackageQuerier{db}
}

type lookupPackageQuerier struct {
	db *sql.DB
}

func (lpq lookupPackageQuerier) Query(packageName, ecosystem string) (PackageVulnerabilities, error) {

	args := []any{packageName}
	query := fmt.Sprint(baseSqlQuery, "WHERE a.package_name = ?")

	if ecosystem != "" {
		args = append(args, ecosystem)
		query = fmt.Sprint(query, " AND a.package_ecosystem = ?")
	}

	stmt, err := lpq.db.Prepare(query)
	if err != nil {
		log.Printf("error while preparing query: %s", err)
		return nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Printf("error while query db!")
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := mapRows(rows)
	if err != nil {
		return nil, err
	}

	return result, nil
}
