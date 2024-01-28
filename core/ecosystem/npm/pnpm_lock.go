package npm

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"regexp"
	"strings"
)

type pnpmLock struct {
	pkgLineRegex     *regexp.Regexp
	pkgVersionsRegex *regexp.Regexp
}

func newPnpmLock() *pnpmLock {
	return &pnpmLock{
		pkgLineRegex:     regexp.MustCompile(`^/([0-9a-z_/@-]+)@(.*):`),
		pkgVersionsRegex: regexp.MustCompile(`([0-9]+.[0-9]+.[0-9]+)`),
	}
}

func (p pnpmLock) ParsePkgLine(line string) (string, string) {
	pkgLine := strings.TrimSpace(line)
	pkgLineParts := p.pkgLineRegex.FindStringSubmatch(pkgLine)
	if len(pkgLineParts) == 3 {
		pkgName := strings.TrimSpace(pkgLineParts[1])
		pkgVersion := strings.TrimSpace(p.pkgVersionsRegex.FindString(pkgLineParts[2]))
		if pkgVersion == "" {
			return "", ""
		}
		return pkgName, pkgVersion
	}
	return "", ""
}

func parsePackagesFromPnpmLockContent(content string) ecosystem.Packages {
	pl := newPnpmLock()
	pkgs := make(ecosystem.Packages, 0)
	lines := strings.Split(content, common.DetectStringLineEnding(content))

	pkgLineReached := false
	for i := range lines {
		if lines[i] == "" {
			continue
		} else if lines[i] == "packages:" {
			pkgLineReached = true
			continue
		}

		if pkgLineReached {
			if strings.HasPrefix(lines[i], "  /") {
				pkgName, pkgVersion := pl.ParsePkgLine(lines[i])
				if pkgName != "" && pkgVersion != "" {
					pkgs = append(pkgs, NewPackage(pkgName, pkgVersion))
				}
			}
		}
	}
	return pkgs
}
