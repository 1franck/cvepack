package npm

import (
	"cvepack/internal/common"
	"cvepack/internal/ecosystem"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func NewProjectFromNodeModules(path string) ecosystem.Project {
	path = filepath.Join(path, NodeModulesFolder)
	return ecosystem.NewProject(path, EcosystemName, scanNodeModules(path))
}

func scanNodeModules(path string) ecosystem.Packages {
	var packages ecosystem.Packages
	var excludedPaths = []string{".bin", ".pnpm"}

	err := filepath.Walk(path,
		func(p string, file os.FileInfo, err error) error {
			if !file.IsDir() || slices.Contains(excludedPaths, file.Name()) {
				return nil
			}

			packageJsonFile := filepath.Join(p, PackageFile)
			if !common.FileExists(packageJsonFile) {
				return nil
			}

			pkg, err := fileToPackageJson(packageJsonFile)
			if err != nil {
				//skip logging error now, because sometimes, it is an intentional error for test purpose
				//log.Println("Error decoding", p, file.Name(), " : ", err)
				return nil
			}

			packages = append(packages, NewPackage(pkg.Name, pkg.Version))
			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return packages
}
