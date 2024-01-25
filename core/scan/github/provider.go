package github

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
)

var (
	branches      = []string{"master", "main"}
	defaultBranch = "master"
)

type Provider struct {
	source ecosystem.Source
	url    *Url
}

func NewProvider(source ecosystem.Source) *Provider {
	return &Provider{source, NewUrl(source)}
}

func (p *Provider) GetPaths(file string) []string {
	paths := make([]string, 0)
	for _, branch := range branches {
		paths = append(paths, p.url.GetFileRawUrl(branch, file))
	}
	return paths
}

func (p *Provider) Source() ecosystem.Source {
	return p.source
}

func (p *Provider) GetFirstExistingPath(file string) string {
	paths := p.GetPaths(file)
	for _, path := range paths {
		if urlExists := common.UrlExists(path, 8); urlExists {
			return path
		}
	}
	return ""
}

// GetPathContent returns the content of path relative to the provider source
func (p *Provider) GetPathContent(file string) (string, error) {
	content, err := common.DownloadUrlContent(file)
	return string(content), err
}
