package scan

import (
	"cvepack/core/ecosystem"
	"cvepack/core/search"
)

type Results []*PackageVulnerabilitiesResult

type PackageVulnerabilitiesResult struct {
	Query           search.PackageVulnerabilityQuery
	Vulnerabilities search.PackageVulnerabilities
}

func (result *Results) LongestPackageName() int {
	var longest int
	for _, value := range *result {
		curLength := value.Query.StringLen()
		if curLength > longest {
			longest = curLength
		}
	}
	return longest
}

func (result *Results) Append(pkg ecosystem.Package, vul search.PackageVulnerabilities) {
	*result = append(*result, &PackageVulnerabilitiesResult{
		Query: search.PackageVulnerabilityQuery{
			Name: pkg.Name(), Version: pkg.Version(), Parent: pkg.Parent(),
		},
		Vulnerabilities: vul,
	})
}

func (result *Results) UniqueResultCount() int {
	unique := make(map[string]bool)
	for _, value := range *result {
		unique[value.Query.ToString()] = true
	}
	return len(unique)
}
