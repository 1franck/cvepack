package search

import "database/sql"

func mapRows(rows *sql.Rows) (PackageVulnerabilities, error) {
	var result PackageVulnerabilities
	for rows.Next() {
		vul, err := mapRow(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, vul)
	}
	return result, nil
}

func mapRow(rows *sql.Rows) (*PackageVulnerability, error) {
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
	err = vul.Parse()
	if err != nil {
		return nil, err
	}

	return &vul, nil
}
