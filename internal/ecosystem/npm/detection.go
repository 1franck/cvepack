package npm

import (
	"github.com/1franck/cvepack/internal/common"
	"path/filepath"
)

func DetectPackageJson(path string) bool {
	if common.FileExists(filepath.Join(path, PackageFile)) {
		return true
	}
	return false
}

func DetectPackageLockJson(path string) bool {
	if common.FileExists(filepath.Join(path, PackageLockFile)) {
		return true
	}
	return false
}

func DetectNodeModules(path string) bool {
	if common.DirectoryExists(filepath.Join(path, NodeModulesFolder)) {
		return true
	}
	return false
}
