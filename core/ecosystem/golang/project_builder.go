package golang

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {

	if DetectGoMod(path) {
		return ecosystem.NewProjectBuilder(
			NewProjectFromGoSum,
			"go.sum detected!")
	}

	return nil
}
