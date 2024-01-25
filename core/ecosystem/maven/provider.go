package maven

import (
	es "cvepack/core/ecosystem"
	"errors"
	"fmt"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	pomXmlPath := provider.GetFirstExistingPath(PomXml)

	if pomXmlPath != "" {
		content, err := es.ProviderPathContent(provider, pomXmlPath)
		if err != nil {
			return nil, err
		}
		return es.NewProject(provider.Source().Value, EcosystemName, parsePackagesFromPomXmlContent(content)), nil
	}

	return nil, errors.New(fmt.Sprintf("no %s project found", EcosystemTitle()))
}
