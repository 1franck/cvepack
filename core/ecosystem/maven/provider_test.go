package maven

import (
	"cvepack/core/scan/files"
	"testing"
)

func Test_NewProjectFromProvider(t *testing.T) {
	provider := files.NewProviderFromPath("./testdata")
	project, err := NewProjectFromProvider(provider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(project.Packages()) != 2 {
		t.Errorf("Expected project to have 2 packages, got %d", len(project.Packages()))
	}
}
