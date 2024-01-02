package scan

import (
	"cvepack/core/ecosystem"
	"fmt"
)

type ScanUrl struct {
	Url      string
	Projects []ecosystem.Project
	Verbose  bool
}

func NewScanUrl(url string) *ScanUrl {
	return &ScanUrl{Url: url}
}

func (scan *ScanUrl) Log(msg string) {
	if scan.Verbose {
		fmt.Println(msg)
	}
}

func (scan *ScanUrl) Run() {

}
