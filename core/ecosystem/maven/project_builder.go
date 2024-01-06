package maven

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {

	if DetectPomXml(path) {
		return ecosystem.NewProjectBuilder(
			NewProjectFromPomXml,
			"pom.xml detected!")
	}

	return nil
}
