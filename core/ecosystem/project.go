package ecosystem

type Project interface {
	Path() string
	Ecosystem() string
	Packages() []Package
}

type defaultProject struct {
	path      string
	ecosystem string
	packages  []Package
}

func NewProject(path, ecosystem string, packages []Package) *defaultProject {
	return &defaultProject{path, ecosystem, packages}
}

func (p *defaultProject) Path() string {
	return p.path
}

func (p *defaultProject) Ecosystem() string {
	return p.ecosystem
}

func (p *defaultProject) Packages() []Package {
	return p.packages
}
