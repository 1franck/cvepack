package npm

import (
	"cvepack/core/common"
	"testing"
)

func Test_PnpmLock_ParsePkgLine(t *testing.T) {
	pnpmLock := newPnpmLock()

	var tests = []struct {
		line               string
		expectedPkgName    string
		expectedPkgVersion string
	}{
		{"/@babel/code-frame@7.5.5:", "@babel/code-frame", "7.5.5"},
		{
			"/svelte-preprocess@5.1.3(postcss@8.4.33)(sass@1.70.0)(svelte@4.2.9)(typescript@5.3.3):",
			"svelte-preprocess",
			"5.1.3",
		},
		{"/send@0.18.0:", "send", "0.18.0"},
		{"/eslint-plugin-unicorn@50.0.1(eslint@8.56.0):", "eslint-plugin-unicorn", "50.0.1"},
		{
			"/@babel/helper-annotate-as-pure@7.16.8(@babel/helper-wrap-function@7.16.7):",
			"@babel/helper-annotate-as-pure",
			"7.16.8",
		},
		{
			"sadfsdfsdfd",
			"",
			"",
		},
		{
			"/send@345----sdf:",
			"",
			"",
		},
	}

	for _, test := range tests {
		pkgName, pkgVersion := pnpmLock.ParsePkgLine(test.line)
		if pkgName != test.expectedPkgName {
			t.Errorf("Expected '%s' pkgName to be '%s', got '%s'", test.line, test.expectedPkgName, pkgName)
		}
		if pkgVersion != test.expectedPkgVersion {
			t.Errorf("Expected pkgVersion to be '%s', got '%s'", test.expectedPkgVersion, pkgVersion)
		}
	}
}

func Test_parsePackagesFromPnpmLockContent(t *testing.T) {
	fileContent, err := common.ReadAllFile("./testdata/pnpm/pnpm-lock.yaml")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	pkgs := parsePackagesFromPnpmLockContent(string(fileContent))
	if len(pkgs) != 939 {
		t.Errorf("Expected to have 939 packages, got %d", len(pkgs))
	}
}
