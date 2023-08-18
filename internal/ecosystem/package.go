package ecosystem

type Package interface {
	Name() string
	Version() string
}
