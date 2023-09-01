package maven

import (
	"cvepack/core/ecosystem"
	"encoding/xml"
	"io"
	"log"
	"os"
	"path/filepath"
)

type pomXml struct {
	XMLName        xml.Name `xml:"project"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	ModelVersion   string   `xml:"modelVersion"`
	GroupId        string   `xml:"groupId"`
	ArtifactId     string   `xml:"artifactId"`
	Version        string   `xml:"version"`
	Dependencies   struct {
		Text       string `xml:",chardata"`
		Dependency []struct {
			Text       string `xml:",chardata"`
			GroupId    string `xml:"groupId"`
			ArtifactId string `xml:"artifactId"`
			Version    string `xml:"version"`
			Scope      string `xml:"scope"`
		} `xml:"dependency"`
	} `xml:"dependencies"`
}

func NewProjectFromPomXml(path string) ecosystem.Project {
	file := filepath.Join(path, PomXml)
	return ecosystem.NewProject(path, EcosystemName, scanPackagesFromPomXml(file))
}

func scanPackagesFromPomXml(file string) ecosystem.Packages {
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

	var project pomXml
	err = xml.Unmarshal(byteValue, &project)
	if err != nil {
		log.Println(err)
		return pkgs
	}

	for _, dep := range project.Dependencies.Dependency {
		name := dep.GroupId + ":" + dep.ArtifactId
		pkgs = append(pkgs, ecosystem.NewDefaultPackage(name, dep.Version, ""))

	}
	return pkgs
}
