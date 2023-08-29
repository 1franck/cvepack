package rubygems

import (
	"cvepack/internal/ecosystem/rubygems"
	"testing"
)

func Test_BuildProjectFromGemFileLock(t *testing.T) {
	project := rubygems.NewProjectFromGemFileLock("./_fixtures")

	if len(project.Packages()) != 46 {
		t.Errorf("Expected project to have 46 packages, got %d", len(project.Packages()))
	}
}
