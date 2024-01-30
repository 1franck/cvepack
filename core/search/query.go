package search

var baseSqlQuery = `
	SELECT a.id, a.vulnerability_id, a.package_ecosystem, a.package_name, ar.e_fixed, ar.e_introduced, ar.e_last_affected, v.summary, v.details, v.aliases, v.database_specific
    FROM affected a
	JOIN vulnerabilities v ON a.vulnerability_id = v.id
    JOIN affected_ranges ar ON a.id = ar.affected_id
`

type PackageVulnerabilityQuery struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Parent  string `json:"parent"`
}

func (q PackageVulnerabilityQuery) ToString() string {
	return q.Name + " " + q.Version
}

func (q PackageVulnerabilityQuery) StringLen() int {
	return len(q.Name + " " + q.Version)
}
