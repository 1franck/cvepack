package cratesio

import (
	es "cvepack/core/ecosystem"
	"errors"
	"fmt"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	cargoFilePath := provider.GetFirstExistingPath(CargoFile)

	// check for Cargo.toml first
	if cargoFilePath != "" {
		cargoLockFilePath := provider.GetFirstExistingPath(CargoLockFile)

		// then check for Cargo.lock
		if cargoLockFilePath != "" {
			content, err := es.ProviderPathContent(provider, cargoLockFilePath)
			if err != nil {
				return nil, err
			}
			return es.NewProject(provider.Source().Value, EcosystemName, parseCargoLockContent(content)), nil
		}

		// otherwise, fall back to Cargo.toml
		// todo: support Cargo.toml
	}

	return nil, errors.New(fmt.Sprintf("no %s project found", EcosystemTitle()))
}
