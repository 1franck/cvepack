package golang

import (
	"github.com/1franck/cvepack/internal/common"
	"path/filepath"
)

func DetectGoMod(path string) bool {
	if common.FileExists(filepath.Join(path, "go.mod")) {
		return true
	}
	return false
}

func DetectGoSum(path string) bool {
	if common.FileExists(filepath.Join(path, "go.sum")) {
		return true
	}
	return false
}
