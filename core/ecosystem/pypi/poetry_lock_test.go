package pypi

import (
	"testing"
)

func Test_BuildProjectFromPoetryLock(t *testing.T) {
	project := NewProjectFromPoetryLock("./testdata")

	if len(project.Packages()) != 74 {
		t.Errorf("Expected project to have 74 packages, got %d", len(project.Packages()))
	}
}
