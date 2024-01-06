package packagist

import (
	"testing"
)

func Test_BuildProjectFromComposerLock(t *testing.T) {
	project := NewProjectFromComposerLock("./testdata")
	if project.Ecosystem() != "Packagist" {
		t.Errorf("Expected project name to be 'Packagist', got %s", project.Ecosystem())
	}

	if len(project.Packages()) != 43 {
		t.Errorf("Expected project to have 43 packages, got %d", len(project.Packages()))
	}
}
