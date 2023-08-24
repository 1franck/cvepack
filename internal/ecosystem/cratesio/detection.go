package cratesio

import (
	"github.com/1franck/cvepack/internal/common"
	"path/filepath"
)

func DetectCargoToml(path string) bool {
	return common.FileExists(filepath.Join(path, CargoFile))
}

func DetectCargoLock(path string) bool {
	return common.FileExists(filepath.Join(path, CargoLockFile))
}
