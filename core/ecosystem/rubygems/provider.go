package rubygems

import (
	es "cvepack/core/ecosystem"
	"errors"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	gemFileLockPath := provider.GetFirstExistingPath(GemFileLock)

	if gemFileLockPath != "" {
		content, err := es.ProviderPathContent(provider, gemFileLockPath)
		if err != nil {
			return nil, err
		}
		return es.NewProject(provider.Source().Value, EcosystemName, parsePackagesGemFileLockContent(content)), nil
	}
	return nil, errors.New("no rubygems project found")
}
