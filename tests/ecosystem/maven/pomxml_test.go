package maven

import (
	"cvepack/core/ecosystem/maven"
	"testing"
)

func Test_NewProjectFromPomXml(t *testing.T) {
	project := maven.NewProjectFromPomXml("./_fixtures")

	if len(project.Packages()) != 2 {
		t.Errorf("Expected project to have 2 packages, got %d", len(project.Packages()))
	}
}
