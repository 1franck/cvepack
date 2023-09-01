package cratesio

import (
	"cvepack/core/ecosystem/nuget"
	"testing"
)

func Test_NewProjectFromSln(t *testing.T) {
	project := nuget.NewProjectFromSln("./_fixtures", "project.sln")

	if len(project.Packages()) != 6 {
		t.Errorf("Expected project to have 6 packages, got %d", len(project.Packages()))
	}
}
