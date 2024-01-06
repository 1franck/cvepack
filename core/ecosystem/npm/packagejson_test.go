package npm

import "testing"

func Test_BuildFromPackageLockJson(t *testing.T) {
	npmProject := NewProjectFromPackageLockJson("./testdata")
	if npmProject.Ecosystem() != "npm" {
		t.Errorf("Expected project name to be 'npm', got %s", npmProject.Ecosystem())
	}

	if len(npmProject.Packages()) != 409 {
		t.Errorf("Expected project to have 409 packages, got %d", len(npmProject.Packages()))
	}
}
