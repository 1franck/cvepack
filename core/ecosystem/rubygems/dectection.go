package rubygems

import (
	"cvepack/core/common"
	"path/filepath"
)

func DetectGemFile(path string) bool {
	return common.FileExists(filepath.Join(path, GemFile))

}

func DetectGemFileLock(path string) bool {
	return common.FileExists(filepath.Join(path, GemFileLock))
}
