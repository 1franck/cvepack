package stats

import (
	"cvepack/core/database"
	"cvepack/core/ecosystem/cratesio"
	"cvepack/core/ecosystem/golang"
	"cvepack/core/ecosystem/maven"
	"cvepack/core/ecosystem/npm"
	"cvepack/core/ecosystem/nuget"
	"cvepack/core/ecosystem/packagist"
	"cvepack/core/ecosystem/pypi"
	"cvepack/core/ecosystem/rubygems"
	"database/sql"
)

var Ecosystems = map[string]string{
	cratesio.EcosystemName:  cratesio.EcosystemLanguage,
	golang.EcosystemName:    golang.EcosystemLanguage,
	maven.EcosystemName:     maven.EcosystemLanguage,
	npm.EcosystemName:       npm.EcosystemLanguage,
	nuget.EcosystemName:     nuget.EcosystemLanguage,
	packagist.EcosystemName: packagist.EcosystemLanguage,
	pypi.EcosystemName:      pypi.EcosystemLanguage,
	rubygems.EcosystemName:  rubygems.EcosystemLanguage,
}

func GetEcosystemsStats(db *sql.DB) (map[string]int, error) {
	result := make(map[string]int)
	for ecosystem := range Ecosystems {
		count, err := database.CountVulnerabilitiesByEcosystem(db, ecosystem)
		if err != nil {
			return result, err
		}
		result[ecosystem] = count
	}
	return result, nil
}
