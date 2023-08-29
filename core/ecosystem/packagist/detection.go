package packagist

import (
	"cvepack/core/common"
	"path/filepath"
)

func DetectComposerJson(path string) bool {
	return common.FileExists(filepath.Join(path, ComposerFile))
}

func DetectComposerLock(path string) bool {
	return common.FileExists(filepath.Join(path, ComposerLockFile))
}
