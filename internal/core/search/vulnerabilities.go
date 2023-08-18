package search

import (
	"fmt"
	"strings"
)

type PackageVulnerabilities []*PackageVulnerability

func (p PackageVulnerabilities) IsEmpty() bool {
	return len(p) < 1
}

func (p PackageVulnerabilities) CountTitle() string {
	vulText := "vulnerability"
	count := len(p)
	if count > 1 {
		vulText = "vulnerabilities"
	}
	return fmt.Sprintf("%d %s", count, vulText)
}

func (p PackageVulnerabilities) SeveritiesSummary() string {
	var result []string
	for level, count := range p.SeveritiesSummaryMap() {
		result = append(result, fmt.Sprintf("%d %s", count, strings.ToLower(level)))
	}
	return strings.Join(result, ", ")
}

func (p PackageVulnerabilities) SeveritiesSummaryMap() map[string]int {
	severities := make(map[string]int, 0)
	for _, v := range p {
		severities[v.SeverityLevel()]++
	}
	return severities
}

func (p PackageVulnerabilities) AliasesSummary() string {
	var result []string
	for _, v := range p {
		aliases := v.AliasesToString()
		if aliases != "" {
			result = append(result, v.AliasesToString())
		}
	}
	return strings.Join(result, ", ")
}
