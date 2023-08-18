package npm

type npmPackage struct {
	name    string
	version string
}

func NewPackage(name string, version string) *npmPackage {
	return &npmPackage{name: name, version: version}
}

func (pkg *npmPackage) Name() string {
	return pkg.name
}

func (pkg *npmPackage) Version() string {
	return pkg.version
}
