package npm

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"regexp"
	"strings"
)

func parsePackagesFromYarnLockContent(content string) ecosystem.Packages {
	pkgs := make(ecosystem.Packages, 0)
	pkgRegex := regexp.MustCompile(`^["]?[^@ ]+`)

	lines := strings.Split(content, common.DetectStringLineEnding(content))

	for i := range lines {
		if lines[i] == "" || strings.HasPrefix(lines[i], "#") || strings.HasPrefix(lines[i], " ") {
			continue
		}

		pkgName := pkgRegex.FindString(lines[i])
		if pkgName == "" {
			continue
		}

		pkgVersionLine := strings.TrimSpace(lines[i+1])
		pkgVersionLineParts := strings.Split(pkgVersionLine, " ")
		if len(pkgVersionLineParts) == 2 {
			pkgs = append(pkgs, NewPackage(pkgName, pkgVersionLineParts[1]))
		}
	}
	return pkgs
}
