package npm

import (
	"log"
	"path/filepath"
)

func NewProjectFromPackageJson(path string) *Project {
	npm := &Project{_path: path}
	packages, err := fileToPackageJson(filepath.Join(path, "package.json"))
	if err == nil {
		for p, ver := range packages.Dependencies {
			npm._packages = append(npm._packages, NewPackage(p, ver))
		}
	}
	if err != nil {
		log.Println(filepath.Join(path, "package.json"))
		log.Println(err)
	}
	return npm
}

func NewProjectFromNodeModules(path string) *Project {
	npm := &Project{_path: path}
	packages := ScanNodeModules(path)
	for _, pkg := range packages {
		npm._packages = append(npm._packages, pkg)
	}
	return npm
}

func NewProjectFromPackageLockJson(path string) *Project {
	npm := &Project{_path: path}
	pkgLock, err := fileToPackageLockJson(filepath.Join(path, "package-lock.json"))
	if err == nil {
		for p, pkg := range pkgLock.Packages {
			if p == "" {
				continue
			}
			p = p[13:]
			npm._packages = append(npm._packages, NewPackage(p, pkg.Version))
		}
	}
	if err != nil {
		log.Println(filepath.Join(path, "package-lock.json"))
		log.Println(err)
	}
	return npm
}
