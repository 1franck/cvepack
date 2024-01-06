package golang

import "testing"

func Test_NewProjectFromGoSum(t *testing.T) {
	project := NewProjectFromGoSum("./testdata")

	if len(project.Packages()) != 35 {
		t.Errorf("Expected project to have 35 packages, got %d", len(project.Packages()))
	}
}
