package scan

import (
	"cvepack/core/ecosystem"
	"cvepack/core/scan/files"
	"cvepack/core/scan/github"
	"fmt"
)

func Inspect(source ecosystem.Source) *ScannerResult {
	s := newScanner()
	return s.Scan(source)
}

type scanner struct {
}

func newScanner() *scanner {
	return &scanner{}
}

// Scan scans a source for projects
// support url and path sources
func (s *scanner) Scan(source ecosystem.Source) *ScannerResult {
	scanResult := NewScannerResult(source)
	defer scanResult.End()

	switch source.Type() {
	case ecosystem.UrlSource:
		scanResult.Projects = s.scanUrl(source)
	case ecosystem.PathSource:
		fileProvider := files.NewProvider(source)
		scanResult.Projects = ProviderScanner(fileProvider)
	}

	return scanResult
}

// scanUrl scan url for projects
func (s *scanner) scanUrl(source ecosystem.Source) []ecosystem.Project {
	fmt.Printf("Url %s\n", source.Value)
	projects := make([]ecosystem.Project, 0)
	if github.DetectGithubRepoUrl(source) {
		provider := github.NewProvider(source)
		return ProviderScanner(provider)
	}
	return projects
}
