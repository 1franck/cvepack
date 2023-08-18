package npm

import (
	"encoding/json"
	"github.com/1franck/cvepack/internal/common"
)

type PackageJson struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type PackageLockJson struct {
	Name            string                         `json:"name"`
	Version         string                         `json:"version"`
	LockfileVersion int                            `json:"lockfileVersion"`
	Requires        bool                           `json:"requires"`
	Packages        map[string]packageLockPackages `json:"packages"`
}

type packageLockPackages struct {
	Version string `json:"version"`
}

func fileToPackageJson(filePath string) (*PackageJson, error) {
	content, err := common.ReadAllFile(filePath)
	if err != nil {
		return nil, err
	}

	var pkg PackageJson
	if err := json.Unmarshal(content, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

func fileToPackageLockJson(filePath string) (*PackageLockJson, error) {
	content, err := common.ReadAllFile(filePath)
	if err != nil {
		return nil, err
	}

	var pkgLock PackageLockJson
	if err := json.Unmarshal(content, &pkgLock); err != nil {
		return nil, err
	}

	return &pkgLock, nil
}

func getPackageJsonVersion(filePath string) (string, error) {
	pkg, err := fileToPackageJson(filePath)
	if err != nil {
		return "", err
	}
	return pkg.Version, nil
}
