package pypi

import (
	"cvepack/core/common"
	"path/filepath"
)

func DetectPyProjectToml(path string) bool {
	return common.FileExists(filepath.Join(path, PyProjectToml))
}

func DetectPoetryLock(path string) bool {
	return common.FileExists(filepath.Join(path, PoetryLock))
}

func DetectPdmLock(path string) bool {
	return common.FileExists(filepath.Join(path, PdmLock))
}
