package github

import "cvepack/core/ecosystem"

func DetectGithubRepoUrl(source ecosystem.Source) bool {
	return NewUrl(source).IsValid
}
