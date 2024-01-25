package nuget

import (
	es "cvepack/core/ecosystem"
	"cvepack/core/scan/files"
	"testing"
)

func Test_NewProjectFromProvider(t *testing.T) {
	provider := files.NewProvider(es.NewPathSource("./testdata"))
	project, err := NewProjectFromProvider(provider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(project.Packages()) != 6 {
		t.Errorf("Expected project to have 6 packages, got %d", len(project.Packages()))
	}
}
