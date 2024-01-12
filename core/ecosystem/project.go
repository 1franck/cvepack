package ecosystem

type Project interface {
	Source() string
	Ecosystem() string
	Packages() []Package
}

type defaultProject struct {
	source    string
	ecosystem string
	packages  []Package
}

func NewProject(source, ecosystem string, packages []Package) *defaultProject {
	return &defaultProject{source, ecosystem, packages}
}

func (p *defaultProject) Source() string {
	return p.source
}

func (p *defaultProject) Ecosystem() string {
	return p.ecosystem
}

func (p *defaultProject) Packages() []Package {
	return p.packages
}

type ProjectBuilder struct {
	Build       func(path string) Project
	Description string
}

func NewProjectBuilder(fn func(path string) Project, desc string) *ProjectBuilder {
	return &ProjectBuilder{fn, desc}
}
