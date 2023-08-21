package scan

import (
	"github.com/1franck/cvepack/internal/core/search"
	"github.com/1franck/cvepack/internal/ecosystem"
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
