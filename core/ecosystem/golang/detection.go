package golang

import (
	"cvepack/core/common"
	"path/filepath"
)

func DetectGoMod(path string) bool {
	return common.FileExists(filepath.Join(path, "go.mod"))

}

func DetectGoSum(path string) bool {
	return common.FileExists(filepath.Join(path, "go.sum"))
}
