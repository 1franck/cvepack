package golang

import (
	"github.com/1franck/cvepack/internal/common"
	"github.com/1franck/cvepack/internal/ecosystem"
	"log"
	"path/filepath"
	"strings"
)

func NewProjectFromGoSum(path string) ecosystem.Project {
	pkgs := ecosystem.Packages{}
	goSumContent, err := common.ReadAllFile(filepath.Join(path, "go.sum"))
	if err != nil {
		log.Println(err)
		return ecosystem.NewProject(path, EcosystemName, pkgs)
	}

	lines := strings.Split(string(goSumContent), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 3 || strings.HasSuffix(parts[1], "/go.mod") {
			continue
		}

		name := parts[0]
		version := parts[1][1:]
		pkgs = append(pkgs, ecosystem.NewDefaultPackage(name, version, ""))
	}

	return ecosystem.NewProject(path, EcosystemName, pkgs)
}
