package ecosystem

type Package interface {
	Name() string
	Version() string
	Parent() string
}

type Packages []Package

type defaultPackage struct {
	name    string
	version string
	parent  string
}

func NewDefaultPackage(name, version, parent string) *defaultPackage {
	return &defaultPackage{name, version, parent}
}

func (pkg *defaultPackage) Name() string {
	return pkg.name
}

func (pkg *defaultPackage) Version() string {
	return pkg.version
}

func (pkg *defaultPackage) Parent() string {
	return pkg.parent
}
