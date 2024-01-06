package packagist

import (
	"cvepack/core/ecosystem"
	"log"
	"path/filepath"
)

type composerLockFile struct {
	Packages    []composerLockPackage `json:"packages"`
	PackagesDev []composerLockPackage `json:"packages-dev"`
}

type composerLockPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewProjectFromComposerLock(path string) ecosystem.Project {
	pkgs := ecosystem.Packages{}
	file := filepath.Join(path, ComposerLockFile)
	composerLock, err := readComposerLockFile(file)

	if err == nil {
		for _, pkg := range composerLock.Packages {
			pkgs = append(pkgs, ecosystem.NewDefaultPackage(pkg.Name, pkg.Version, ""))
		}
		for _, pkg := range composerLock.PackagesDev {
			pkgs = append(pkgs, ecosystem.NewDefaultPackage(pkg.Name, pkg.Version, ""))
		}
	} else {
		log.Printf("Error while loadding file %s : %s", file, err)
	}

	return ecosystem.NewProject(path, EcosystemName, pkgs)
}
