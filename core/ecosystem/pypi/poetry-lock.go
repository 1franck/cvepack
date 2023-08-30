package pypi

import (
	"cvepack/core/ecosystem"
	"log"
	"path/filepath"
)

func NewProjectFromPoetryLock(path string) ecosystem.Project {
	file := filepath.Join(path, PoetryLock)
	pkgs, err := parseLockContent(file)

	if err != nil {
		log.Println(err)
	}

	return ecosystem.NewProject(path, EcosystemName, pkgs)
}
