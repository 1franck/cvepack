package pypi

import (
	"cvepack/internal/ecosystem"
	"strings"
)

func parseLockContent(content []byte) ecosystem.Packages {
	pkgs := ecosystem.Packages{}
	lines := strings.Split(string(content), "\n")

	for i, line := range lines {
		if line == "[[package]]" {
			name := strings.Replace(lines[i+1][7:], "\"", "", -1)
			version := strings.Replace(lines[i+2][9:], "\"", "", -1)
			pkgs = append(pkgs, ecosystem.NewDefaultPackage(name, strings.TrimSpace(version), ""))
		}
	}
	return pkgs
}
