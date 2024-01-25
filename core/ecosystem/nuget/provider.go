package nuget

import (
	es "cvepack/core/ecosystem"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func NewProjectFromProvider(provider es.Provider) (es.Project, error) {
	if provider.Source().Type() == es.UrlSource {
		return nil, errors.New(fmt.Sprintf("NuGet url scanning not supported yet"))
	} else if provider.Source().Type() == es.UnknownSource {
		return nil, errors.New(fmt.Sprintf("Unknown source"))
	}

	files, err := os.ReadDir(provider.Source().Value)
	if err != nil {
		return nil, err
	}

	var slnFile string

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sln") {
			slnFile = file.Name()
			break
		}
	}

	if slnFile != "" {
		log.Printf("Found solution file: %s", slnFile)
		pkgs := es.Packages{}
		file := filepath.Join(provider.Source().Value, slnFile)
		csprojFiles := scanCsprojFromSln(file)

		if len(csprojFiles) == 0 {
			return nil, errors.New("no csproj file(s) found")
		}

		for _, csproj := range csprojFiles {
			pkgs.Append(scanPackagesFromCsProjXml(
				filepath.Join(provider.Source().Value, csproj))...)
		}

		return es.NewProject(provider.Source().Value, EcosystemName, pkgs), nil
	}
	return nil, errors.New("no .sln file found")

}
