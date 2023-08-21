package npm

import (
	"encoding/json"
	"github.com/1franck/cvepack/internal/common"
)

func fileToPackageLockJson(filePath string) (*PackageLockJsonFile, error) {
	content, err := common.ReadAllFile(filePath)
	if err != nil {
		return nil, err
	}

	var pkgLock PackageLockJsonFile
	if err := json.Unmarshal(content, &pkgLock); err != nil {
		return nil, err
	}

	return &pkgLock, nil
}
