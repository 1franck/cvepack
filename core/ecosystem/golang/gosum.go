package golang

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"log"
	"path/filepath"
	"strings"
)

func NewProjectFromGoSum(path string) ecosystem.Project {
	pkgs := ecosystem.Packages{}
	file := filepath.Join(path, "go.sum")
	goSumContent, err := common.ReadAllFile(file)
	if err != nil {
		log.Println(err)
		return ecosystem.NewProject(path, EcosystemName, pkgs)
	}

	lines := strings.Split(string(goSumContent), common.DetectLineEnding(file))

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
		pkgs.Append(ecosystem.NewDefaultPackage(name, version, ""))
	}

	return ecosystem.NewProject(path, EcosystemName, pkgs)
}
