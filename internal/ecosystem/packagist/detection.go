package packagist

import (
	"cvepack/internal/common"
	"path/filepath"
)

func DetectComposerJson(path string) bool {
	return common.FileExists(filepath.Join(path, ComposerFile))
}

func DetectComposerLock(path string) bool {
	return common.FileExists(filepath.Join(path, ComposerLockFile))
}
