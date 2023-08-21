package npm

type PackageLockJsonFile struct {
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
