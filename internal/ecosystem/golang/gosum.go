package golang

import (
	"github.com/1franck/cvepack/internal/common"
	"log"
	"path/filepath"
	"strings"
)

func NewProjectFromGoSum(path string) *Project {
	project := &Project{_path: path}
	gosumContent, err := common.ReadAllFile(filepath.Join(path, "go.sum"))
	if err != nil {
		log.Println(err)
		return project
	}

	lines := strings.Split(string(gosumContent), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 3 || strings.HasSuffix(parts[1], "/go.mod") {
			continue
		}

		pkgName := parts[0]
		pkgVersion := parts[1][1:]

		pkg := NewPackage(pkgName, pkgVersion)
		project._packages = append(project._packages, pkg)
	}

	return project
}
