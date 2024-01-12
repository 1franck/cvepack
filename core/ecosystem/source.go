package ecosystem

type SourceType string

const (
	UrlSource     SourceType = "url"
	PathSource    SourceType = "path"
	UnknownSource SourceType = "unknown"
)

type Source struct {
	Name  string
	_type SourceType
}

func NewSource(name string, sourceType SourceType) Source {
	return Source{Name: name, _type: sourceType}
}

func (s *Source) Type() SourceType {
	return s._type
}

func normalizeSourceType(t string) SourceType {
	switch SourceType(t) {
	case UrlSource:
		return UrlSource
	case PathSource:
		return PathSource
	default:
		return UnknownSource
	}
}
