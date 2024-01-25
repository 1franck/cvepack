package ecosystem

import (
	"cvepack/core/common"
	"errors"
	"fmt"
	"strings"
)

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

func NewPathSource(value string) Source {
	return NewSource(value, PathSource)
}

func NewUrlSource(value string) Source {
	return NewSource(value, UrlSource)
}

func (s Source) Type() SourceType {
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

func ValidateSource(source Source) error {
	if strings.TrimSpace(source.Value) == "" {
		return errors.New("source is empty")
	}
	if source.Type() == UnknownSource {
		return errors.New(fmt.Sprintf("unknown source: %s", source.Value))
	}

	if source.Type() == PathSource {
		if err := common.ValidateDirectory(source.Value); err != nil {
			return err
		}
	}
	return nil
}

func ErrorUnknownSourceType(source Source) error {
	return errors.New(fmt.Sprintf("unknown source for %s", source.Value))
}
