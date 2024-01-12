package analysis

import (
	"cvepack/core/scan"
	"github.com/axllent/semver"
)

func LongestPackageName(result *scan.PackagesVulnerabilitiesResult) int {
	var longest int
	for _, value := range *result {
		curLength := value.Query.StringLen()
		if curLength > longest {
			longest = curLength
		}
	}
	return longest
}

func ShortestPackageName(result *scan.PackagesVulnerabilitiesResult) int {
	var shortest int
	for _, value := range *result {
		curLength := value.Query.StringLen()
		if curLength < shortest {
			shortest = curLength
		}
	}
	return shortest
}

func VersionToUpdate(pkgVul *scan.PackageVulnerabilitiesResult) string {
	versions := make([]string, 0)
	for _, vul := range pkgVul.Vulnerabilities {
		if hasFix, fix := vul.HasFix(); hasFix {
			versions = append(versions, fix)
		}
	}

	if len(versions) < 1 {
		return ""
	}

	versions = semver.SortMax(versions)
	return versions[0]
}
