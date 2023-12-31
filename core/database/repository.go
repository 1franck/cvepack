package database

import (
	"cvepack/core/osv"
	"cvepack/core/sqlite"
	"database/sql"
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

func CountVulnerabilitiesByEcosystem(db *sql.DB, ecosystem string) (int, error) {
	var count int
	err := db.QueryRow(countVulnerabilityByEcosystemQuery, ecosystem).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
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
