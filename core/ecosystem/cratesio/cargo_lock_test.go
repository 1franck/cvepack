package cratesio

import (
	"testing"
)

func Test_BuildProjectFromCargoLock(t *testing.T) {
	project := NewProjectFromCargoLock("./testdata")

	if len(project.Packages()) != 180 {
		t.Errorf("Expected project to have 180 packages, got %d", len(project.Packages()))
	}
}
