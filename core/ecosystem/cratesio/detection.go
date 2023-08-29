package cratesio

import (
	"cvepack/core/common"
	"path/filepath"
)

func DetectCargoToml(path string) bool {
	return common.FileExists(filepath.Join(path, CargoFile))
}

func DetectCargoLock(path string) bool {
	return common.FileExists(filepath.Join(path, CargoLockFile))
}
