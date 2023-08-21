package search

type PackageVulnerabilityQuery struct {
	Name    string
	Version string
	Parent  string
}

func (q PackageVulnerabilityQuery) ToString() string {
	return q.Name + " " + q.Version
}

func (q PackageVulnerabilityQuery) StringLen() int {
	return len(q.Name + " " + q.Version)
}
