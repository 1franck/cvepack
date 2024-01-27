package npm

import (
	"cvepack/core/common"
	"testing"
)

func Test_parsePackagesFromYarnLockContent_with_v1(t *testing.T) {
	fileContent, err := common.ReadAllFile("./testdata/yarnv1/yarn.lock")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	pkgs := parsePackagesFromYarnLockContent(string(fileContent))
	if len(pkgs) != 436 {
		t.Errorf("Expected to have 436 packages, got %d", len(pkgs))
	}
}

func Test_parsePackagesFromYarnLockContent_with_v2plus(t *testing.T) {
	fileContent, err := common.ReadAllFile("./testdata/yarnv2+/yarn.lock")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	pkgs := parsePackagesFromYarnLockContent(string(fileContent))
	if len(pkgs) != 2062 {
		t.Errorf("Expected to have 2062 packages, got %d", len(pkgs))
	}
}
