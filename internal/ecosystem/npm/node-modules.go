package npm

import (
	"github.com/1franck/cvepack/internal/ecosystem"
	"log"
	"os"
	"path/filepath"
	"slices"
)

var nodeModulesExcludedPaths = []string{".bin", ".pnpm"}

func ScanNodeModules(path string) []ecosystem.Package {

	var packages []ecosystem.Package
	// Read the contents of the folder
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the files and subdirectories in the folder
	for _, file := range files {
		if file.IsDir() && !excludedPathInNodeModules(file.Name()) {
			version, err := getPackageJsonVersion(filepath.Join(path, file.Name(), "package.json"))
			if err != nil {
				log.Println(err)
				continue
			}
			packages = append(packages, NewPackage(file.Name(), version))
		}
	}
	return packages
}

func excludedPathInNodeModules(path string) bool {
	if slices.Contains(nodeModulesExcludedPaths, path) {
		return true
	}
	return false
}
