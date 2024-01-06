package pypi

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {

	if DetectPyProjectToml(path) {
		if DetectPoetryLock(path) {
			return ecosystem.NewProjectBuilder(
				NewProjectFromPoetryLock,
				"poetry.lock detected!")
		} else if DetectPdmLock(path) {
			return ecosystem.NewProjectBuilder(
				NewProjectFromPdmLock,
				"pdm.lock detected!")
		}
	}

	return nil
}
