package npm

import "encoding/json"

type packageLockJson struct {
	Name            string                        `json:"name"`
	Version         string                        `json:"version"`
	LockfileVersion int                           `json:"lockfileVersion"`
	Requires        bool                          `json:"requires"`
	Packages        map[string]packageLockPackage `json:"packages"`
}

type packageLockPackage struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Dependencies map[string]string `json:"dependencies"`
}

func stringToPackageLockJson(content string) (*packageLockJson, error) {
	var pkgLock packageLockJson
	if err := json.Unmarshal([]byte(content), &pkgLock); err != nil {
		return nil, err
	}

	return &pkgLock, nil
}
