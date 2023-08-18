package ecosystem

type Ecosystem interface {
	Name() string
	Packages() []Package
}
