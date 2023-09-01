package nuget

import "encoding/xml"

type CsProject struct {
	XMLName   xml.Name `xml:"Project"`
	Text      string   `xml:",chardata"`
	Sdk       string   `xml:"Sdk,attr"`
	ItemGroup []struct {
		Text             string `xml:",chardata"`
		PackageReference []struct {
			Text    string `xml:",chardata"`
			Include string `xml:"Include,attr"`
			Version string `xml:"Version,attr"`
		} `xml:"PackageReference"`
	} `xml:"ItemGroup"`
}
