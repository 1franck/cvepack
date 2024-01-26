package pypi

import (
	es "cvepack/core/ecosystem"
	"cvepack/core/scan/files"
	"testing"
)

func Test_PyPi_NewProjectFromProvider(t *testing.T) {
	provider := files.NewProvider(es.NewSource("./testdata", es.PathSource))
	project, err := NewProjectFromProvider(provider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(project.Packages()) != 74 {
		t.Errorf("Expected project to have 74 packages, got %d", len(project.Packages()))
	}
}
