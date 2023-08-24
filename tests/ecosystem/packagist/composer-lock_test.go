package packagist

import (
	"github.com/1franck/cvepack/internal/ecosystem/packagist"
	"testing"
)

func Test_BuildProjectFromComposerLock(t *testing.T) {
	project := packagist.NewProjectFromComposerLock("./_fixtures")
	if project.Ecosystem() != "Packagist" {
		t.Errorf("Expected project name to be 'Packagist', got %s", project.Ecosystem())
	}

	if len(project.Packages()) != 43 {
		t.Errorf("Expected project to have 43 packages, got %d", len(project.Packages()))
	}
}
