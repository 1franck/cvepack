package npm

import (
	"log"
	"path/filepath"
)

func NewProjectFromPackageLockJson(path string) *Project {
	npm := &Project{_path: path}
	pkgLock, err := fileToPackageLockJson(filepath.Join(path, "package-lock.json"))

	if err == nil {
		for pkgKey, pkg := range pkgLock.Packages {
			if pkgKey == "" {
				// reference to the package itself or an installed node_modules package, skip
				continue
			} else if pkg.Name != "" && pkg.Name != pkgKey {
				// reference to a local package
				pkgKey = pkg.Name
			}
			npm._packages = append(npm._packages, NewPackage(pkgKey, pkg.Version))
		}
	} else {
		log.Println(filepath.Join(path, "package-lock.json"))
		log.Println(err)
	}

	return npm
}

func findParents(
	parents map[string][]string,
	pkgLock map[string]packageLockPackage,
	pkgName string, pkgVersion string) map[string][]string {
	for key, pkgDef := range pkgLock {
		for depName, depVersion := range pkgDef.Dependencies {
			if depName == pkgName && depVersion == pkgVersion {
				if _, ok := parents[key]; !ok {
					parents[key] = []string{}
				}
				parents[key] = append(parents[key], pkgName)

			}
		}
	}
	return parents
}
