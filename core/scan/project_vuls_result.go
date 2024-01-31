package scan

import (
	"cvepack/core/ecosystem"
	"encoding/json"
	"path/filepath"
	"time"
)

type ProjectVulnerabilitiesResult struct {
	Source     string                        `json:"source"`
	Ecosystem  string                        `json:"ecosystem"`
	Date       time.Time                     `json:"date"`
	ScanResult PackagesVulnerabilitiesResult `json:"results"`
}

type ProjectsVulnerabilitiesResult []*ProjectVulnerabilitiesResult

func (r *ProjectsVulnerabilitiesResult) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *ProjectsVulnerabilitiesResult) Add(project ecosystem.Project, pkgsVulsResult *PackagesVulnerabilitiesResult) {
	sourceAbs, err := filepath.Abs(project.Source())
	if err != nil {
		sourceAbs = project.Source()
	}
	*r = append(*r, &ProjectVulnerabilitiesResult{
		Source:     sourceAbs,
		Ecosystem:  project.Ecosystem(),
		Date:       time.Now(),
		ScanResult: *pkgsVulsResult,
	})
}
