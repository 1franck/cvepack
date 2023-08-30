package pypi

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"strings"
)

func parseLockContent(file string) (ecosystem.Packages, error) {
	pkgs := ecosystem.Packages{}

	lockContent, err := common.ReadAllFile(file)

	if err != nil {
		return pkgs, err
	}

	lineEnding := common.DetectLineEnding(file)
	lines := strings.Split(string(lockContent), lineEnding)

	for i, line := range lines {
		if line == "[[package]]" {
			name := strings.Replace(lines[i+1][7:], "\"", "", -1)
			version := strings.Replace(lines[i+2][9:], "\"", "", -1)
			pkgs = append(pkgs, ecosystem.NewDefaultPackage(name, strings.TrimSpace(version), ""))
		}
	}

	return pkgs, nil
}
