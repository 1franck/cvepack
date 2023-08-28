package rubygems

import (
	"fmt"
	"github.com/1franck/cvepack/internal/common"
	"github.com/1franck/cvepack/internal/ecosystem"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var packageRegex = regexp.MustCompile(`([a-zA-Z_-]+)\s+\((\d+\.\d+\.\d+)\)`)

type gemFilePackage struct {
	Name    string
	Version string
}

func NewProjectFromGemFileLock(path string) ecosystem.Project {
	pkgs := ecosystem.Packages{}
	gemLockContent, err := common.ReadAllFile(filepath.Join(path, GemFileLock))
	if err != nil {
		log.Println(err)
		return ecosystem.NewProject(path, EcosystemName, pkgs)
	}

	lines := strings.Split(string(gemLockContent), "\n")
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

	return ecosystem.NewProject(path, EcosystemName, pkgs)
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
