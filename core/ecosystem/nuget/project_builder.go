package nuget

import "cvepack/core/ecosystem"

func ProjectBuilder(path string) *ecosystem.ProjectBuilder {
	slnFile := DetectSln(path)
	if slnFile != "" {
		return ecosystem.NewProjectBuilder(
			func(path string) ecosystem.Project {
				return NewProjectFromSln(path, slnFile)
			},
			".sln file detected!")
	}
	return nil
}
