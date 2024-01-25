package npm

import (
	"cvepack/core/common"
	"encoding/json"
)

//func fileToPackageLockJson(filePath string) (*packageLockJson, error) {
//	content, err := common.ReadAllFile(filePath)
//	if err != nil {
//		return nil, err
//	}
//
//	var pkgLock packageLockJson
//	if err := json.Unmarshal(content, &pkgLock); err != nil {
//		return nil, err
//	}
//
//	return &pkgLock, nil
//}

func fileToPackageJson(filePath string) (*packageJson, error) {
	content, err := common.ReadAllFile(filePath)
	if err != nil {
		return nil, err
	}

	var pkg packageJson
	if err := json.Unmarshal(content, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}
