package search

import (
	"database/sql"
	"github.com/axllent/semver"
	"log"
)

var query = `
    SELECT a.id, a.vulnerability_id, a.package_ecosystem, a.package_name, ar.e_fixed, ar.e_introduced, ar.e_last_affected, v.summary, v.details, v.aliases, v.database_specific
    FROM affected a
	JOIN vulnerabilities v ON a.vulnerability_id = v.id
    JOIN affected_ranges ar ON a.id = ar.affected_id
    WHERE a.package_ecosystem = ? AND a.package_name = ?
`

func PackageVulnerabilityQuerier(db *sql.DB) *packageVulnerabilityQuerier {
	return &packageVulnerabilityQuerier{db}
}

type packageVulnerabilityQuerier struct {
	db *sql.DB
}

func (pvq packageVulnerabilityQuerier) Query(ecosystem, packageName string, packageVersion string) (PackageVulnerabilities, error) {
	stmt, err := pvq.db.Prepare(query)
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
		if semver.Compare(*pkg.VersionFixed, pkgVersion) == -1 { // 1 mean vulnerability VersionFixed > current pkgVersion
			return false
		}
		return true
	} else if pkg.VersionLastAffected != nil {
		if semver.Compare(*pkg.VersionLastAffected, pkgVersion) == -1 { // 1 mean vulnerability VersionLastAffected > current pkgVersion
			return false
		}
		return true
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
