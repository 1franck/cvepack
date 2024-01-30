package scan

import (
	"cvepack/core/ecosystem"
	"cvepack/core/search"
	"encoding/json"
)

type PackagesVulnerabilitiesResult []*PackageVulnerabilitiesResult

type PackageVulnerabilitiesResult struct {
	Query           search.PackageVulnerabilityQuery `json:"Package"`
	Vulnerabilities search.PackageVulnerabilities    `json:"Vulnerabilities"`
}

func (result *PackagesVulnerabilitiesResult) LongestPackageName() int {
	var longest int
	for _, value := range *result {
		curLength := value.Query.StringLen()
		if curLength > longest {
			longest = curLength
		}
	}
	return longest
}

func (result *PackagesVulnerabilitiesResult) Append(pkg ecosystem.Package, vul search.PackageVulnerabilities) {
	*result = append(*result, &PackageVulnerabilitiesResult{
		Query: search.PackageVulnerabilityQuery{
			Name: pkg.Name(), Version: pkg.Version(), Parent: pkg.Parent(),
		},
		Vulnerabilities: vul,
	})
}

func (result *PackagesVulnerabilitiesResult) UniqueResultCount() int {
	unique := make(map[string]bool)
	for _, value := range *result {
		unique[value.Query.ToString()] = true
	}
	return len(unique)
}

func (result *PackagesVulnerabilitiesResult) ToJson() ([]byte, error) {
	return json.Marshal(result)
}
