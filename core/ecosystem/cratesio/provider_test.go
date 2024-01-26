package cratesio

import (
	"cvepack/core/scan/files"
	"testing"
)

func Test_CratesIo_NewProjectFromProvider(t *testing.T) {
	provider := files.NewProviderFromPath("./testdata")
	project, err := NewProjectFromProvider(provider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(project.Packages()) != 180 {
		t.Errorf("Expected project to have 180 packages, got %d", len(project.Packages()))
	}
}
