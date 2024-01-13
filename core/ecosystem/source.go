package ecosystem

type SourceType string

const (
	UrlSource     SourceType = "url"
	PathSource    SourceType = "path"
	UnknownSource SourceType = "unknown"
)

type Source struct {
	Value string
	_type SourceType
}

func NewSource(value string, sourceType SourceType) Source {
	return Source{Value: value, _type: sourceType}
}

func (s *Source) Type() SourceType {
	return s._type
}

func StringToSourceType(t string) SourceType {
	switch SourceType(t) {
	case UrlSource:
		return UrlSource
	case PathSource:
		return PathSource
	default:
		return UnknownSource
	}
}
