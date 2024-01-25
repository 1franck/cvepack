package files

import (
	"cvepack/core/common"
	"cvepack/core/ecosystem"
	"path/filepath"
)

type Provider struct {
	source ecosystem.Source
}

func NewProvider(source ecosystem.Source) *Provider {
	return &Provider{source}
}

func NewProviderFromPath(path string) *Provider {
	return NewProvider(ecosystem.NewPathSource(path))
}

func (p *Provider) GetPaths(file string) []string {
	fp := filepath.Join(p.source.Value, file)
	return []string{fp}
}

func (p *Provider) Source() ecosystem.Source {
	return p.source
}

func (p *Provider) GetFirstExistingPath(file string) string {
	paths := p.GetPaths(file)
	for _, path := range paths {
		if common.FileExists(path) {
			return path
		}
	}
	return ""
}
