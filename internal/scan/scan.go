package scan

import (
	"fmt"
	"github.com/1franck/cvepack/internal/ecosystem"
	"github.com/1franck/cvepack/internal/ecosystem/npm"
	"path/filepath"
)

type Scan struct {
	Path       string
	Ecosystems []ecosystem.Ecosystem
	Verbose    bool
}

func NewScan(path string) *Scan {
	return &Scan{Path: path}
}

func (scan *Scan) Run() {
	if npm.DetectPackageJson(scan.Path) {
		if npm.DetectPackageLockJson(scan.Path) {
			scan.Log("package-lock.json detected ...")
			scan.Ecosystems = append(scan.Ecosystems, npm.NewProjectFromPackageLockJson(scan.Path))
		} else if npm.DetectNodeModules(scan.Path) {
			scan.Log("/node_modules detected ...")
			scan.Ecosystems = append(scan.Ecosystems, npm.NewProjectFromNodeModules(filepath.Join(scan.Path, "node_modules")))
		}
	}
}

func (scan *Scan) Log(msg string) {
	if scan.Verbose {
		fmt.Println(msg)
	}
}
