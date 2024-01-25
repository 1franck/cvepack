package scan

import (
	es "cvepack/core/ecosystem"
	"cvepack/core/ecosystem/cratesio"
	"cvepack/core/ecosystem/golang"
	"cvepack/core/ecosystem/maven"
	"cvepack/core/ecosystem/npm"
	"cvepack/core/ecosystem/nuget"
	"cvepack/core/ecosystem/packagist"
	"cvepack/core/ecosystem/pypi"
	"cvepack/core/ecosystem/rubygems"
)

func ProviderScanner(provider es.Provider) []es.Project {
	projects := make([]es.Project, 0)

	cratesioProject, err := cratesio.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, cratesioProject)
	}

	golangProject, err := golang.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, golangProject)
	}

	mavenProject, err := maven.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, mavenProject)
	}

	npmProject, err := npm.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, npmProject)
	}

	nugetProject, err := nuget.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, nugetProject)
	}

	packagistProject, err := packagist.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, packagistProject)
	}

	pypiProject, err := pypi.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, pypiProject)
	}

	rubyGems, err := rubygems.NewProjectFromProvider(provider)
	if err == nil {
		projects = append(projects, rubyGems)
	}

	return projects
}
