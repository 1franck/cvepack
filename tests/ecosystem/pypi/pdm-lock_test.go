package pypi

import (
	"cvepack/core/ecosystem/pypi"
	"testing"
)

func Test_BuildProjectFromPdmLock(t *testing.T) {
	project := pypi.NewProjectFromPdmLock("./_fixtures")

	if len(project.Packages()) != 47 {
		t.Errorf("Expected project to have 47 packages, got %d", len(project.Packages()))
	}
}
