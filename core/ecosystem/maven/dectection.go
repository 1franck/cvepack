package maven

import (
	"cvepack/core/common"
	"path/filepath"
)

func DetectPomXml(path string) bool {
	return common.FileExists(filepath.Join(path, PomXml))
}
