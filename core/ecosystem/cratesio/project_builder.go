package cratesio

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {

	if DetectCargoToml(path) {
		return ecosystem.NewProjectBuilder(
			NewProjectFromCargoLock,
			"Cargo.lock detected!")
	}

	return nil
}
