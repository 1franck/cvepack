package packagist

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {

	if DetectComposerJson(path) {
		return ecosystem.NewProjectBuilder(
			NewProjectFromComposerLock,
			"composer.json detected!")
	}

	return nil
}
