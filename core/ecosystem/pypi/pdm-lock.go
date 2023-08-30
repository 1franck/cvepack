package pypi

import (
	"cvepack/core/ecosystem"
	"log"
	"path/filepath"
)

func NewProjectFromPdmLock(path string) ecosystem.Project {
	file := filepath.Join(path, PdmLock)

	pkgs, err := parseLockContent(file)

	if err != nil {
		log.Println(err)
	}

	return ecosystem.NewProject(path, EcosystemName, pkgs)
}
