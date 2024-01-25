package npm

import (
	es "cvepack/core/ecosystem"
	"errors"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	packageJsonPath := provider.GetFirstExistingPath(PackageFile)

	// check for package.json first
	if packageJsonPath != "" {
		packageJsonLockPath := provider.GetFirstExistingPath(PackageLockFile)

		// then check for package-lock.json
		if packageJsonLockPath != "" {
			content, err := es.ProviderPathContent(provider, packageJsonLockPath)
			if err != nil {
				return nil, err
			}
			return StringPackageLockJsonToProject(content, packageJsonLockPath), nil
		}

		// otherwise, check for node_modules
		nodeModulesFolder := provider.GetFirstExistingPath(NodeModulesFolder)
		if nodeModulesFolder != "" {
			return es.NewProject(nodeModulesFolder, EcosystemName, scanNodeModules(nodeModulesFolder)), nil
		}

		// otherwise, check for yarn.lock
		// TODO support yarn.lock

		// otherwise, use package.json
		// TODO support package.json
	}

	return nil, errors.New("no package.json found")
}
