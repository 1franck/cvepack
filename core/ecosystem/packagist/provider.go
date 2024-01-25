package packagist

import (
	es "cvepack/core/ecosystem"
	"encoding/json"
	"errors"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	composerPath := provider.GetFirstExistingPath(ComposerFile)

	if composerPath != "" {
		composerLockPath := provider.GetFirstExistingPath(ComposerLockFile)

		if composerLockPath != "" {
			content, err := es.ProviderPathContent(provider, composerLockPath)
			if err != nil {
				return nil, err
			}

			composerLock, err := stringToComposerLockJson(content)
			if err != nil {
				return nil, err
			}

			pkgs := parsePackagesFromComposerLockFile(*composerLock)
			return es.NewProject(composerLockPath, EcosystemName, pkgs), nil
		}

		return nil, errors.New("no composer.lock found")
	}

	return nil, errors.New("no composer.json found")
}

func stringToComposerLockJson(content string) (*composerLockFile, error) {
	var composerLockFile composerLockFile
	if err := json.Unmarshal([]byte(content), &composerLockFile); err != nil {
		return nil, err
	}
	return &composerLockFile, nil
}
