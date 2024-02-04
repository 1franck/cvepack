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
	"fmt"
	"sort"
)

var (
	Ecosystems = map[string]string{
		cratesio.EcosystemName:  cratesio.EcosystemLanguage,
		golang.EcosystemName:    golang.EcosystemLanguage,
		maven.EcosystemName:     maven.EcosystemLanguage,
		npm.EcosystemName:       npm.EcosystemLanguage,
		nuget.EcosystemName:     nuget.EcosystemLanguage,
		packagist.EcosystemName: packagist.EcosystemLanguage,
		pypi.EcosystemName:      pypi.EcosystemLanguage,
		rubygems.EcosystemName:  rubygems.EcosystemLanguage,
	}
	ecosystemKeys []string
)

func init() {
	ecosystemKeys = make([]string, 0)
	for k, _ := range Ecosystems {
		ecosystemKeys = append(ecosystemKeys, k)
	}
	sort.Strings(ecosystemKeys)
}

type EcosystemStats struct {
	Name     string
	Language string
	Count    int
}

func NewEcosystemStats(name string, count int) EcosystemStats {
	language := Ecosystems[name]
	return EcosystemStats{
		Name:     name,
		Language: language,
		Count:    count,
	}
}

func (e EcosystemStats) Percentage(total int) float64 {
	return float64(e.Count) / float64(total) * 100
}

func (e EcosystemStats) GetTitle() string {
	if e.Language != e.Name {
		return fmt.Sprintf("%s (%s)", e.Name, e.Language)
	}
	return e.Language
}

func GetEcosystemsStats(db *sql.DB) ([]EcosystemStats, error) {
	result := make([]EcosystemStats, 0)
	for _, ecosystem := range ecosystemKeys {
		count, err := database.CountVulnerabilitiesByEcosystem(db, ecosystem)
		if err != nil {
			return result, err
		}
		result = append(result, NewEcosystemStats(ecosystem, count))
	}
	return result, nil
}
