package maven

import (
	"testing"
)

func Test_NewProjectFromPomXml(t *testing.T) {
	project := NewProjectFromPomXml("./testdata")

	if len(project.Packages()) != 2 {
		t.Errorf("Expected project to have 2 packages, got %d", len(project.Packages()))
	}
}
