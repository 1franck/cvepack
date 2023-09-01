package nuget

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"encoding/xml"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var slnCsprojRegex = regexp.MustCompile(`"([^"]+\.csproj)"`)

func NewProjectFromSln(path string, slnFile string) ecosystem.Project {
	file := filepath.Join(path, slnFile)

	pkgs := ecosystem.Packages{}
	for _, csproj := range scanCsprojFromSln(file) {
		pkgs.Append(scanPackagesFromCsProjXml(filepath.Join(path, csproj))...)
	}

	return ecosystem.NewProject(path, EcosystemName, pkgs)
}

func scanCsprojFromSln(file string) []string {
	content, err := common.ReadAllFile(file)
	if err != nil {
		log.Println(err)
		return []string{}
	}

	lineEnding := common.DetectLineEnding(file)
	lines := strings.Split(string(content), lineEnding)
	var csprojPaths []string

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "Project") {
			matches := slnCsprojRegex.FindStringSubmatch(trimmedLine)
			if len(matches) > 1 {
				proj := strings.Replace(matches[1], "\\", "/", -1)
				csprojPaths = append(csprojPaths, proj)
			}
		}
	}

	return csprojPaths
}

func scanPackagesFromCsProjXml(file string) ecosystem.Packages {
	pkgs := ecosystem.Packages{}

	xmlFile, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return pkgs
	}

	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			panic(err)
		}
	}(xmlFile)

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		log.Println(err)
		return pkgs
	}

	var project CsProject
	err = xml.Unmarshal(byteValue, &project)
	if err != nil {
		log.Println(err)
		return pkgs
	}

	for _, itemGroup := range project.ItemGroup {
		for _, ref := range itemGroup.PackageReference {
			pkgs = append(pkgs, ecosystem.NewDefaultPackage(ref.Include, ref.Version, ""))
		}
	}
	return pkgs
}
