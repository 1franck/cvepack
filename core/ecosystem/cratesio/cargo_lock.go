package cratesio

import (
	"cvepack/core/common"
	es "cvepack/core/ecosystem"
	"strings"
)

func parseCargoLockContent(content string) es.Packages {
	pkgs := es.Packages{}
	lines := strings.Split(content, common.DetectStringLineEnding(content))

	for i, line := range lines {
		if line == "[[package]]" {
			name := strings.Replace(lines[i+1][7:], "\"", "", -1)
			version := strings.Replace(lines[i+2][9:], "\"", "", -1)
			pkgs = append(pkgs, es.NewDefaultPackage(name, strings.TrimSpace(version), ""))
		}
	}
	return pkgs
}
