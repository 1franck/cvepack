package nuget

import (
	"testing"
)

func Test_NewProjectFromSln(t *testing.T) {
	project := NewProjectFromSln("./testdata", "project.sln")

	if len(project.Packages()) != 6 {
		t.Errorf("Expected project to have 6 packages, got %d", len(project.Packages()))
	}
}
