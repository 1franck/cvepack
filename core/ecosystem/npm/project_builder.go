package npm

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {

	if DetectPackageJson(path) {
		if DetectPackageLockJson(path) {
			return ecosystem.NewProjectBuilder(
				NewProjectFromPackageLockJson,
				"package-lock.json detected!")
		} else if DetectNodeModules(path) {
			return ecosystem.NewProjectBuilder(
				NewProjectFromNodeModules,
				"node_modules detected!")
		}
	}

	return nil
}
