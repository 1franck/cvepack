package npm

import (
	"github.com/1franck/cvepack/internal/ecosystem"
)

type Project struct {
	_path     string
	_packages []ecosystem.Package
}

func NewProject(path string, packages []ecosystem.Package) *Project {
	return &Project{_path: path, _packages: packages}
}

func (npm *Project) Name() string {
	return "npm"
}

func (npm *Project) Packages() []ecosystem.Package {
	return npm._packages
}
