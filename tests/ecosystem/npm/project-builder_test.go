package npm

import (
	"github.com/1franck/cvepack/internal/ecosystem/npm"
	"testing"
)

func Test_BuildFromPackageLockJson(t *testing.T) {
	npmProject := npm.NewProjectFromPackageLockJson("./_fixtures/package-lock.json")
	if npmProject.Name() != "npm" {
		t.Errorf("Expected project name to be 'npm', got %s", npmProject.Name())
	}
}
