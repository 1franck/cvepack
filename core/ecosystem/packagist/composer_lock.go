package packagist

import (
	es "cvepack/core/ecosystem"
)

type composerLockFile struct {
	Packages    []composerLockPackage `json:"packages"`
	PackagesDev []composerLockPackage `json:"packages-dev"`
}

type composerLockPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func parsePackagesFromComposerLockFile(composerLock composerLockFile) es.Packages {
	pkgs := es.Packages{}
	for _, pkg := range composerLock.Packages {
		pkgs = append(pkgs, es.NewDefaultPackage(pkg.Name, pkg.Version, ""))
	}
	for _, pkg := range composerLock.PackagesDev {
		pkgs = append(pkgs, es.NewDefaultPackage(pkg.Name, pkg.Version, ""))
	}
	return pkgs
}
