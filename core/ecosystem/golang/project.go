package golang

import "cvepack/core/ecosystem"

func NewProject(source ecosystem.Source, packages []ecosystem.Package) ecosystem.Project {
	return ecosystem.NewProject(source.Value, EcosystemName, packages)
}
