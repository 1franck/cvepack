package pypi

import (
	"testing"
)

func Test_parseLockContent_with_poetryLock(t *testing.T) {
	file := "./testdata/poetry.lock"
	pkgs, err := parseLockContent(file)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(pkgs) != 74 {
		t.Errorf("Expected 74 packages, got %d", len(pkgs))
	}
}

func Test_parseLockContent_with_pdmLock(t *testing.T) {
	file := "./testdata/pdm.lock"
	pkgs, err := parseLockContent(file)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(pkgs) != 47 {
		t.Errorf("Expected 47 packages, got %d", len(pkgs))
	}
}
