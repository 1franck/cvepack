package npm

import (
	"cvepack/core/ecosystem"
	"log"
	"path/filepath"
)

type packageJson struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type packageLockJson struct {
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

func NewProjectFromPackageLockJson(path string) ecosystem.Project {
	file := filepath.Join(path, PackageLockFile)
	pkgLock, err := fileToPackageLockJson(file)

	// todo: return error
	if err != nil {
		log.Printf("Error while loadding file %s : %s", file, err)
		return ecosystem.NewProject(path, EcosystemName, ecosystem.Packages{})
	}

	return PackageLockJsonToProject(pkgLock, path)
}

// PackageLockJsonToProject
// transform packageLockJson to Project
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
