package rubygems

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"fmt"
	"regexp"
	"strings"
)

var packageRegex = regexp.MustCompile(`([a-zA-Z_-]+)\s+\((\d+\.\d+\.\d+)\)`)

type gemFilePackage struct {
	Name    string
	Version string
}

func parsePackageString(text string) (gemFilePackage, error) {
	match := packageRegex.FindStringSubmatch(text)

	if match != nil {
		return gemFilePackage{
			Name:    match[1],
			Version: match[2],
		}, nil
	}
	return gemFilePackage{}, fmt.Errorf("no match found")
}

func parsePackagesGemFileLockContent(content string) ecosystem.Packages {
	pkgs := ecosystem.Packages{}
	lineEnding := common.DetectStringLineEnding(content)
	lines := strings.Split(content, lineEnding)
	specsSection := false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if !specsSection && line != "specs:" {
			continue
		} else if !specsSection && line == "specs:" {
			specsSection = true
			continue
		} else if specsSection && line == "" {
			specsSection = false
			continue
		}

		if specsSection {
			p, err := parsePackageString(line)
			if err == nil {
				pkgs = append(pkgs, ecosystem.NewDefaultPackage(p.Name, p.Version, ""))
			}
		}
	}

	return pkgs
}
