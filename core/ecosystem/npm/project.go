package npm

import (
	"cvepack/core/ecosystem"
	"log"
)

func PackageLockJsonToProject(pkgLock *packageLockJson, source string) ecosystem.Project {
	pkgs := ecosystem.Packages{}

	for pkgKey, pkg := range pkgLock.Packages {
		if pkgKey == "" {
			// reference to the package itself or an installed node_modules package, skip
			continue
		} else if pkg.Name != "" && pkg.Name != pkgKey {
			// reference to a local package
			pkgKey = pkg.Name
		}
		pkgs = append(pkgs, NewPackage(pkgKey, pkg.Version))
	}

	return ecosystem.NewProject(source, EcosystemName, pkgs)
}

func StringPackageLockJsonToProject(pkgJson string, source string) ecosystem.Project {
	pkg, err := stringToPackageLockJson(pkgJson)
	if err != nil {
		log.Printf("Error while loadding file %s : %s", pkgJson, err)
		return ecosystem.NewProject(source, EcosystemName, ecosystem.Packages{})
	}

	return PackageLockJsonToProject(pkg, source)
}
