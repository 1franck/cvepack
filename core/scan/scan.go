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

type Scan struct {
	Path     string
	Projects []ecosystem.Project
	Verbose  bool
}

func NewScan(path string) *Scan {
	return &Scan{Path: path}
}

func (scan *Scan) Log(msg string) {
	if scan.Verbose {
		fmt.Println(msg)
	}
}

func (scan *Scan) Run() {
	waitGroup := sync.WaitGroup{}

	builders := make([]ecosystem.ProjectBuilder, 0)

	npmProjectBuilder := npm.ProjectBuilder(scan.Path)
	if npmProjectBuilder != nil {
		builders = append(builders, *npmProjectBuilder)
	}

	golangProjectBuilder := golang.ProjectBuilder(scan.Path)
	if golangProjectBuilder != nil {
		builders = append(builders, *golangProjectBuilder)
	}

	packagistProjectBuilder := packagist.ProjectBuilder(scan.Path)
	if packagistProjectBuilder != nil {
		builders = append(builders, *packagistProjectBuilder)
	}

	cratesioBuilder := cratesio.ProjectBuilder(scan.Path)
	if cratesioBuilder != nil {
		builders = append(builders, *cratesioBuilder)
	}

	rubygemsBuilder := rubygems.ProjectBuilder(scan.Path)
	if rubygemsBuilder != nil {
		builders = append(builders, *rubygemsBuilder)
	}

	pypiBuilder := pypi.ProjectBuilder(scan.Path)
	if pypiBuilder != nil {
		builders = append(builders, *pypiBuilder)
	}

	nugetBuilder := nuget.ProjectBuilder(scan.Path)
	if nugetBuilder != nil {
		builders = append(builders, *nugetBuilder)
	}

	mavenBuilder := maven.ProjectBuilder(scan.Path)
	if mavenBuilder != nil {
		builders = append(builders, *mavenBuilder)
	}

	for _, builder := range builders {
		waitGroup.Add(1)
		b := builder
		go func() {
			scan.Log(b.Description)
			scan.Projects = append(scan.Projects, b.Build(scan.Path))
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
