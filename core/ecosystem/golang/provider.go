package golang

import (
	es "cvepack/core/ecosystem"
	"errors"
	"fmt"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	goSumPath := provider.GetFirstExistingPath(GoSum)

	if goSumPath != "" {
		content, err := es.ProviderPathContent(provider, goSumPath)
		if err != nil {
			return nil, err
		}
		return es.NewProject(
			provider.Source().Value,
			EcosystemName,
			parsePackagesFromGoSumContent(content)), nil
	}

	return nil, errors.New(fmt.Sprintf("no %s found", GoSum))
}
