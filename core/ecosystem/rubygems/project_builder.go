package rubygems

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {

	if DetectGemFileLock(path) {
		return ecosystem.NewProjectBuilder(
			NewProjectFromGemFileLock,
			"Gemfile.lock detected!")
	}

	return nil
}
