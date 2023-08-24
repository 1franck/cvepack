package golang

type goPackage struct {
	name    string
	version string
	parent  string
}

func NewPackage(name string, version string) *goPackage {
	return &goPackage{name, version, ""}
}

func (pkg *goPackage) Name() string {
	return pkg.name
}

func (pkg *goPackage) Version() string {
	return pkg.version
}

func (pkg *goPackage) Parent() string {
	return pkg.parent
}
