package cratesio

import (
	"cvepack/core/ecosystem/cratesio"
	"testing"
)

func Test_BuildProjectFromCargoLock(t *testing.T) {
	project := cratesio.NewProjectFromCargoLock("./_fixtures")

	if len(project.Packages()) != 180 {
		t.Errorf("Expected project to have 180 packages, got %d", len(project.Packages()))
	}
}
