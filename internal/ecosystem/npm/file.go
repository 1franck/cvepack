package npm

import (
	"cvepack/internal/common"
	"encoding/json"
)

func fileToPackageLockJson(filePath string) (*packageLockJsonFile, error) {
	content, err := common.ReadAllFile(filePath)
	if err != nil {
		return nil, err
	}

	var pkgLock packageLockJsonFile
	if err := json.Unmarshal(content, &pkgLock); err != nil {
		return nil, err
	}

	return &pkgLock, nil
}

func fileToPackageJson(filePath string) (*packageJsonFile, error) {
	content, err := common.ReadAllFile(filePath)
	if err != nil {
		return nil, err
	}

	var pkg packageJsonFile
	if err := json.Unmarshal(content, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}
