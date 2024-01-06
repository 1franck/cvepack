package rubygems

import (
	"testing"
)

func Test_BuildProjectFromGemFileLock(t *testing.T) {
	project := NewProjectFromGemFileLock("./testdata")

	if len(project.Packages()) != 46 {
		t.Errorf("Expected project to have 46 packages, got %d", len(project.Packages()))
	}
}
