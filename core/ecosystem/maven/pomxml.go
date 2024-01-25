package maven

import (
	"cvepack/core/ecosystem"
	"encoding/xml"
	"log"
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

func parsePackagesFromPomXmlContent(content string) ecosystem.Packages {
	pkgs := ecosystem.Packages{}
	var project pomXml
	err := xml.Unmarshal([]byte(content), &project)
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
