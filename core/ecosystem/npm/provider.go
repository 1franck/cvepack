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
		yarnLockPath := provider.GetFirstExistingPath(YarnLockFile)
		if yarnLockPath != "" {
			content, err := es.ProviderPathContent(provider, yarnLockPath)
			if err != nil {
				return nil, err
			}
			return es.NewProject(yarnLockPath, EcosystemName, parsePackagesFromYarnLockContent(content)), nil
		}

		// otherwise, check form pnpm-lock.yaml
		pnpmLockPath := provider.GetFirstExistingPath(PnpmLockFile)
		if pnpmLockPath != "" {
			content, err := es.ProviderPathContent(provider, pnpmLockPath)
			if err != nil {
				return nil, err
			}
			return es.NewProject(pnpmLockPath, EcosystemName, parsePackagesFromPnpmLockContent(content)), nil
		}

		// otherwise, fallback package.json
		// TODO support package.json
	}

	return nil, errors.New("no package.json found")
}
