package core

import (
	"database/sql"
	"github.com/1franck/cvepack/internal/osv"
	"github.com/1franck/cvepack/internal/sqlite"
	"log"
)

func CountVulnerabilities(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(countVulnerabilityQuery).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func SearchPackage(db *sql.DB, ecosystem string, packageName string) ([]interface{}, error) {
	stmt, err := db.Prepare(searchAffectedPackageQuery)
	if err != nil {
		log.Printf("error while preparing query: %s", err)
		return nil, err
	}
	result, err := stmt.Query(ecosystem, packageName)
	if err != nil {
		log.Printf("error while query db!")
		return nil, err
	}
	var rows []interface{}
	err = result.Scan(&rows)
	if err != nil {
		log.Printf("error while scanning results!")
		return nil, err
	}
	return rows, nil
}

func InsertVulnerability(db *sql.DB, v osv.Osv) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	prepareVulnerabilityStmt, err := db.Prepare(insertVulnerabilityQuery)
	if err != nil {
		return err
	}
	defer func(prepareVulnerabilityStmt *sql.Stmt) {
		err := prepareVulnerabilityStmt.Close()
		if err != nil {
			panic(err)
		}
	}(prepareVulnerabilityStmt)

	_, err = sqlite.ExecPrepare(tx.Stmt(prepareVulnerabilityStmt),
		v.ID,
		v.Modified,
		v.Published,
		v.Withdrawn,
		v.AliasesJson(),
		v.RelatedJson(),
		v.Summary,
		v.Details,
		v.SeverityJson(),
		v.ReferencesJson(),
		v.CreditsJson(),
		v.DatabaseSpecificJson())

	if err != nil {
		return err
	}

	prepareAffectedStmt, err := db.Prepare(insertAffectedQuery)
	if err != nil {
		return err
	}
	defer prepareAffectedStmt.Close()

	prepareAffectedRangesStmt, err := db.Prepare(insertAffectedRangesQuery)
	if err != nil {
		return err
	}
	defer prepareAffectedRangesStmt.Close()

	for _, affected := range v.Affected {

		affectedRow, err := sqlite.ExecPrepare(tx.Stmt(prepareAffectedStmt),
			v.ID,
			affected.Package.Ecosystem,
			affected.Package.Name,
			affected.Package.Purl,
			affected.SeverityJson(),
			affected.VersionsJson(),
			affected.EcosystemSpecificJson(),
			affected.DatabaseSpecificJson())

		if err != nil {
			return err
		}

		affectedId, err := affectedRow.LastInsertId()
		if err != nil {
			return err
		}

		for _, affectedRange := range affected.Ranges {

			var eIntroduced *string
			var eFixed *string
			var eLastAffected *string
			var eLimit *string

			for _, affectedRangeEvent := range affectedRange.Events {
				if affectedRangeEvent.Introduced != nil {
					eIntroduced = affectedRangeEvent.Introduced
				}
				if affectedRangeEvent.Fixed != nil {
					eFixed = affectedRangeEvent.Fixed
				}
				if affectedRangeEvent.LastAffected != nil {
					eLastAffected = affectedRangeEvent.LastAffected
				}
				if affectedRangeEvent.Limit != nil {
					eLimit = affectedRangeEvent.Limit
				}
			}

			_, err := sqlite.ExecPrepare(tx.Stmt(prepareAffectedRangesStmt),
				v.ID,
				affectedId,
				affectedRange.Type,
				affectedRange.Repo,
				affectedRange.DatabaseSpecificJson(),
				eIntroduced,
				eFixed,
				eLastAffected,
				eLimit)

			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
