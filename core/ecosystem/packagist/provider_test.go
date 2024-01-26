package packagist

import (
	es "cvepack/core/ecosystem"
	"cvepack/core/scan/files"
	"testing"
)

func Test_Packagist_NewProjectFromProvider(t *testing.T) {
	provider := files.NewProvider(es.NewPathSource("./testdata"))
	project, err := NewProjectFromProvider(provider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(project.Packages()) != 43 {
		t.Errorf("Expected project to have 43 packages, got %d", len(project.Packages()))
	}
}
