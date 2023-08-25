package cratesio

import (
	"github.com/1franck/cvepack/internal/common"
	"github.com/1franck/cvepack/internal/ecosystem"
	"log"
	"path/filepath"
	"strings"
)

func NewProjectFromCargoLock(path string) ecosystem.Project {
	pkgs := ecosystem.Packages{}
	file := filepath.Join(path, CargoLockFile)
	cargoLockContent, err := common.ReadAllFile(file)

	if err != nil {
		log.Println(err)
		return ecosystem.NewProject(path, EcosystemName, pkgs)
	}

	lines := strings.Split(string(cargoLockContent), "\n")

	for i, line := range lines {
		if line == "[[package]]" {
			name := strings.Replace(lines[i+1][7:], "\"", "", -1)
			version := strings.Replace(lines[i+2][9:], "\"", "", -1)
			pkgs = append(pkgs, ecosystem.NewDefaultPackage(name, strings.TrimSpace(version), ""))
		}
	}

	return ecosystem.NewProject(path, EcosystemName, pkgs)
}
