package scan

import (
	"cvepack/core/ecosystem"
	"cvepack/core/ecosystem/cratesio"
	"cvepack/core/ecosystem/golang"
	"cvepack/core/ecosystem/maven"
	"cvepack/core/ecosystem/npm"
	"cvepack/core/ecosystem/nuget"
	"cvepack/core/ecosystem/packagist"
	"cvepack/core/ecosystem/pypi"
	"cvepack/core/ecosystem/rubygems"
	"fmt"
	"sync"
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

func (s *scanner) Scan(source ecosystem.Source) *ScannerResult {
	scanResult := NewScannerResult(source)
	defer scanResult.End()
	switch source.Type() {

	case ecosystem.UrlSource:
		scanResult.Projects = s.scanUrl(source)
	case ecosystem.PathSource:
		scanResult.Projects = s.scanPath(source)
	}

	return scanResult
}

func (s *scanner) scanPath(source ecosystem.Source) []ecosystem.Project {
	path := source.Name
	waitGroup := sync.WaitGroup{}
	projects := make([]ecosystem.Project, 0)
	builders := make([]ecosystem.ProjectBuilder, 0)

	npmProjectBuilder := npm.ProjectBuilder(path)
	if npmProjectBuilder != nil {
		builders = append(builders, *npmProjectBuilder)
	}

	golangProjectBuilder := golang.ProjectBuilder(path)
	if golangProjectBuilder != nil {
		builders = append(builders, *golangProjectBuilder)
	}

	packagistProjectBuilder := packagist.ProjectBuilder(path)
	if packagistProjectBuilder != nil {
		builders = append(builders, *packagistProjectBuilder)
	}

	cratesioBuilder := cratesio.ProjectBuilder(path)
	if cratesioBuilder != nil {
		builders = append(builders, *cratesioBuilder)
	}

	rubygemsBuilder := rubygems.ProjectBuilder(path)
	if rubygemsBuilder != nil {
		builders = append(builders, *rubygemsBuilder)
	}

	pypiBuilder := pypi.ProjectBuilder(path)
	if pypiBuilder != nil {
		builders = append(builders, *pypiBuilder)
	}

	nugetBuilder := nuget.ProjectBuilder(path)
	if nugetBuilder != nil {
		builders = append(builders, *nugetBuilder)
	}

	mavenBuilder := maven.ProjectBuilder(path)
	if mavenBuilder != nil {
		builders = append(builders, *mavenBuilder)
	}

	for _, builder := range builders {
		waitGroup.Add(1)
		b := builder
		go func() {
			projects = append(projects, b.Build(path))
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	return projects
}

func (s *scanner) scanUrl(source ecosystem.Source) []ecosystem.Project {
	fmt.Printf("Scanning url %s", source.Name)
	return []ecosystem.Project{}
}
