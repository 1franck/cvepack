package npm

import (
	"github.com/1franck/cvepack/internal/common"
	"path/filepath"
)

func DetectPackageJson(path string) bool {
	return common.FileExists(filepath.Join(path, PackageFile))
}

func DetectPackageLockJson(path string) bool {
	return common.FileExists(filepath.Join(path, PackageLockFile))
}

func DetectNodeModules(path string) bool {
	return common.DirectoryExists(filepath.Join(path, NodeModulesFolder))
}
