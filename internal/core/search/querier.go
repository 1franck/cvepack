package search

import (
	"database/sql"
	"fmt"
	"github.com/1franck/cvepack/internal/core"
	"log"
)

func PackageVulnerabilityQuerier(db *sql.DB) *packageVulnerabilityQuerier {
	sqlQuery := fmt.Sprint(baseSqlQuery, "WHERE a.package_ecosystem = ? AND a.package_name = ?")
	return &packageVulnerabilityQuerier{db, sqlQuery}
}

type packageVulnerabilityQuerier struct {
	db       *sql.DB
	sqlQuery string
}

func (pvq packageVulnerabilityQuerier) Query(ecosystem, packageName, packageVersion string) (PackageVulnerabilities, error) {
	stmt, err := pvq.db.Prepare(pvq.sqlQuery)
	if err != nil {
		log.Printf("error while preparing query: %s", err)
		return nil, err
	}
	rows, err := stmt.Query(ecosystem, packageName)
	if err != nil {
		log.Printf("error while query db!")
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := pvq.mapRows(rows)
	if err != nil {
		return nil, err
	}

	filteredResult := pvq.removeInactive(result, packageVersion)
	filteredResult = pvq.removeDuplicate(filteredResult)

	//sort.SliceStable(filteredResult, func(i, j int) bool {
	//	return filteredResult[i].PackageName < filteredResult[j].PackageName
	//})

	return filteredResult, nil
}

func (pvq packageVulnerabilityQuerier) mapRows(rows *sql.Rows) (PackageVulnerabilities, error) {
	var result PackageVulnerabilities
	for rows.Next() {
		vul, err := pvq.mapRow(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, vul)
	}
	return result, nil
}

func (pvq packageVulnerabilityQuerier) mapRow(rows *sql.Rows) (*PackageVulnerability, error) {
	var vul PackageVulnerability
	err := rows.Scan(
		&vul.Id,
		&vul.VulnerabilityId,
		&vul.PackageEcosystem,
		&vul.PackageName,
		&vul.VersionFixed,
		&vul.VersionIntroduced,
		&vul.VersionLastAffected,
		&vul.Summary,
		&vul.Details,
		&vul.Aliases,
		&vul.DatabaseSpecific)
	if err != nil {
		return nil, err
	}
	return &vul, nil
}

func (pvq packageVulnerabilityQuerier) isActive(pkg PackageVulnerability, pkgVersion string) bool {
	if pkg.VersionFixed != nil {
		return core.IsVersionAffectedByFixedVersion(
			pkgVersion,
			pkg.VersionIntroduced,
			*pkg.VersionFixed)

	} else if pkg.VersionLastAffected != nil {
		return core.IsVersionInRange(
			pkgVersion,
			pkg.VersionIntroduced,
			*pkg.VersionLastAffected)
	}

	return true
}

func (pvq packageVulnerabilityQuerier) removeInactive(vul PackageVulnerabilities, pkgVersion string) PackageVulnerabilities {
	var result PackageVulnerabilities
	for _, v := range vul {
		if pvq.isActive(*v, pkgVersion) {
			result = append(result, v)
		}
	}
	return result
}

func (pvq packageVulnerabilityQuerier) removeDuplicate(vul PackageVulnerabilities) PackageVulnerabilities {
	var result PackageVulnerabilities
	for _, v := range vul {
		exists := false
		for _, r := range result {
			if v.Id == r.Id {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, v)
		}
	}
	return result
}
