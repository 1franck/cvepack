package pypi

import (
	es "cvepack/core/ecosystem"
	"errors"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	pyProjectTomlPath := provider.GetFirstExistingPath(PyProjectToml)

	if pyProjectTomlPath != "" {
		poetryLockPath := provider.GetFirstExistingPath(PoetryLock)

		if poetryLockPath != "" {
			content, err := es.ProviderPathContent(provider, poetryLockPath)
			if err != nil {
				return nil, err
			}
			return es.NewProject(provider.Source().Value, EcosystemName, parsePackagesLockContent(content)), nil
		}

		pdmLockPath := provider.GetFirstExistingPath(PdmLock)
		if pdmLockPath != "" {
			content, err := es.ProviderPathContent(provider, pdmLockPath)
			if err != nil {
				return nil, err
			}
			return es.NewProject(provider.Source().Value, EcosystemName, parsePackagesLockContent(content)), nil
		}
	}

	return nil, errors.New("no pypi project found")
}
