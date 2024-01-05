package stats

import (
	"cvepack/core/database"
	"database/sql"
)

func GetTotalVulnerabilities(db *sql.DB) int {
	total, err := database.CountVulnerabilities(db)
	if err != nil {
		return 0
	}
	return total
}
