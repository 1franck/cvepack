package scan

import (
	es "cvepack/core/ecosystem"
	"cvepack/core/scan/files"
	"cvepack/core/scan/github"
)

func Inspect(source es.Source) *ScannerResult {
	scanResult := NewScannerResult(source)
	defer scanResult.End()

	switch source.Type() {
	case es.UrlSource:
		scanResult.Projects = inspectUrl(source)
	case es.PathSource:
		fileProvider := files.NewProvider(source)
		scanResult.Projects = ProviderScanner(fileProvider)
	}

	return scanResult
}

func inspectUrl(source es.Source) []es.Project {
	projects := make([]es.Project, 0)
	if github.DetectGithubRepoUrl(source) {
		provider := github.NewProvider(source)
		return ProviderScanner(provider)
	}
	return projects
}
