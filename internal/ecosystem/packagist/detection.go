package packagist

import (
	"github.com/1franck/cvepack/internal/common"
	"path/filepath"
)

func DetectComposerJson(path string) bool {
	if common.FileExists(filepath.Join(path, ComposerFile)) {
		return true
	}
	return false
}

func DetectComposerLock(path string) bool {
	if common.FileExists(filepath.Join(path, ComposerLockFile)) {
		return true
	}
	return false
}
