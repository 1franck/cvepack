package pypi

import (
	"cvepack/internal/common"
	"cvepack/internal/ecosystem"
	"log"
	"path/filepath"
)

func NewProjectFromPoetryLock(path string) ecosystem.Project {
	file := filepath.Join(path, PoetryLock)
	lockContent, err := common.ReadAllFile(file)

	if err != nil {
		log.Println(err)
		return ecosystem.NewProject(path, EcosystemName, ecosystem.Packages{})
	}

	return ecosystem.NewProject(path, EcosystemName, parseLockContent(lockContent))
}
