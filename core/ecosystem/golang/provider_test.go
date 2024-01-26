package golang

import (
	"cvepack/core/scan/files"
	"testing"
)

func Test_Golang_NewProjectFromProvider(t *testing.T) {
	provider := files.NewProviderFromPath("./testdata")
	project, err := NewProjectFromProvider(provider)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(project.Packages()) != 35 {
		t.Errorf("Expected project to have 35 packages, got %d", len(project.Packages()))
	}
}
