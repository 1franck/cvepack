package github

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"regexp"
)

var (
	urlTpl = "https://raw.githubusercontent.com/{repo}/{branch}/{filepath}"
)

type Url struct {
	Source  ecosystem.Source
	Repo    string
	IsValid bool
}

func NewUrl(source ecosystem.Source) *Url {
	gh := &Url{Source: source}
	validateUrl(gh)
	return gh
}

func (gh *Url) GetFileRawUrl(branch, filepath string) string {
	return common.ReplacePlaceholders(urlTpl, map[string]string{
		"repo":     gh.Repo,
		"branch":   branch,
		"filepath": filepath,
	})
}

func (gh *Url) GeneratePath(filepath string) string {
	branch := "master"
	return common.ReplacePlaceholders("{repo}/{branch}/{filepath}", map[string]string{
		"repo":     gh.Repo,
		"branch":   branch,
		"filepath": filepath,
	})
}

func validateUrl(gh *Url) {
	regexStr := `^(?:https?:\/\/)?github\.com\/([a-zA-Z0-9\-_]+\/[a-zA-Z0-9\-_]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(gh.Source.Value)
	if len(matches) == 0 {
		return
	}

	gh.Repo = matches[1]
	gh.IsValid = true
}
