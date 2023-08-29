package pypi

import (
	"cvepack/internal/ecosystem/pypi"
	"testing"
)

func Test_BuildProjectFromPoetryLock(t *testing.T) {
	project := pypi.NewProjectFromPoetryLock("./_fixtures")

	if len(project.Packages()) != 74 {
		t.Errorf("Expected project to have 74 packages, got %d", len(project.Packages()))
	}
}
