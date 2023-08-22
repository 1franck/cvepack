package npm

import (
	"log"
	"path/filepath"
)

type packageJsonFile struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type packageLockJsonFile struct {
	Name            string                        `json:"name"`
	Version         string                        `json:"version"`
	LockfileVersion int                           `json:"lockfileVersion"`
	Requires        bool                          `json:"requires"`
	Packages        map[string]packageLockPackage `json:"packages"`
}

type packageLockPackage struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Dependencies map[string]string `json:"dependencies"`
}

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
