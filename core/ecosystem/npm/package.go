package npm

import (
	"strings"
)

type npmPackage struct {
	rawName string
	name    string
	version string
	parent  string
}

func NewPackage(name string, version string) *npmPackage {
	finalName := resolvePackageName(name)
	parent := resolvePackageParentName(name)
	return &npmPackage{name, finalName, version, parent}
}

func resolvePackageName(name string) string {
	if strings.HasPrefix(name, "node_modules/") {
		name = strings.TrimPrefix(name, "node_modules/")
	}
	if strings.Contains(name, "/node_modules/") {
		parts := strings.Split(name, "/node_modules/")
		return parts[len(parts)-1]
	}
	return name
}

func resolvePackageParentName(name string) string {
	if strings.HasPrefix(name, "node_modules/") {
		name = strings.TrimPrefix(name, "node_modules/")
	}
	if strings.Contains(name, "/node_modules/") {
		parts := strings.Split(name, "/node_modules/")
		return parts[len(parts)-2]
	}
	return ""
}

func (pkg *npmPackage) Name() string {
	return pkg.name
}

func (pkg *npmPackage) Version() string {
	return pkg.version
}

func (pkg *npmPackage) Parent() string {
	return pkg.parent
}
