package ecosystem

import "cvepack/core/common"

type Provider interface {
	// GetPaths returns a list of paths to the file in the provider's source
	GetPaths(file string) []string

	// GetFirstExistingPath retrieve the path of the first existing file in the provider's source
	GetFirstExistingPath(file string) string

	// Source returns the provider's source
	Source() Source
}

func ProviderPathContent(provider Provider, file string) (string, error) {
	switch provider.Source().Type() {
	case PathSource:
		content, err := common.ReadAllFile(file)
		return string(content), err
	case UrlSource:
		content, err := common.DownloadUrlContent(file)
		return string(content), err
	default:
		return "", ErrorUnknownSourceType(provider.Source())
	}
}
