package pypi

import (
	"testing"
)

func Test_BuildProjectFromPdmLock(t *testing.T) {
	project := NewProjectFromPdmLock("./testdata")

	if len(project.Packages()) != 47 {
		t.Errorf("Expected project to have 47 packages, got %d", len(project.Packages()))
	}
}
