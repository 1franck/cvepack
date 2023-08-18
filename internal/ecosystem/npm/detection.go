package npm

import (
	"github.com/1franck/cvepack/internal/common"
	"path/filepath"
)

func DetectPackageJson(path string) bool {
	if common.FileExists(filepath.Join(path, "package.json")) {
		return true
	}
	return false
}

func DetectPackageLockJson(path string) bool {
	if common.FileExists(filepath.Join(path, "package-lock.json")) {
		return true
	}
	return false
}

func DetectNodeModules(path string) bool {
	if common.DirectoryExists(filepath.Join(path, "node_modules")) {
		return true
	}
	return false
}

func DetectYarnLock(path string) bool {
	if common.FileExists(filepath.Join(path, "yarn.lock")) {
		return true
	}
	return false
}
