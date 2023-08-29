package npm

import (
	"cvepack/core/ecosystem/npm"
	"testing"
)

func Test_BuildFromPackageLockJson(t *testing.T) {
	npmProject := npm.NewProjectFromPackageLockJson("./_fixtures")
	if npmProject.Ecosystem() != "npm" {
		t.Errorf("Expected project name to be 'npm', got %s", npmProject.Ecosystem())
	}

	if len(npmProject.Packages()) != 409 {
		t.Errorf("Expected project to have 409 packages, got %d", len(npmProject.Packages()))
	}
}
