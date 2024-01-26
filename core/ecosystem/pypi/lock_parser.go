package pypi

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"strings"
)

func parsePackagesLockContent(content string) ecosystem.Packages {
	lineEnding := common.DetectStringLineEnding(content)
	lines := strings.Split(content, lineEnding)
	pkgs := ecosystem.Packages{}

	for i := range lines {
		if lines[i] == "[[package]]" {
			name := strings.Replace(lines[i+1][7:], "\"", "", -1)
			version := strings.Replace(lines[i+2][9:], "\"", "", -1)
			pkgs = append(pkgs, ecosystem.NewDefaultPackage(name, strings.TrimSpace(version), ""))
		}
	}

	return pkgs
}
