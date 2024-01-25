package golang

import (
	"cvepack/core/common"
	es "cvepack/core/ecosystem"
	"strings"
)

func parsePackagesFromGoSumContent(content string) es.Packages {
	pkgs := es.Packages{}
	lines := strings.Split(content, common.DetectStringLineEnding(content))

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 3 || strings.HasSuffix(parts[1], "/go.mod") {
			continue
		}

		name := parts[0]
		version := parts[1][1:]
		pkgs.Append(es.NewDefaultPackage(name, version, ""))
	}
	return pkgs
}
