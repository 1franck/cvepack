package scan

import (
	"fmt"
	"github.com/1franck/cvepack/internal/ecosystem"
	"github.com/1franck/cvepack/internal/ecosystem/golang"
	"github.com/1franck/cvepack/internal/ecosystem/npm"
	"sync"
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
	waitGroup := sync.WaitGroup{}

	if npm.DetectPackageJson(scan.Path) {
		if npm.DetectPackageLockJson(scan.Path) {
			scan.Log("package-lock.json detected!")
			waitGroup.Add(1)
			go func() {
				scan.Ecosystems = append(scan.Ecosystems, npm.NewProjectFromPackageLockJson(scan.Path))
				waitGroup.Done()
			}()
		} else if npm.DetectNodeModules(scan.Path) {
			scan.Log("node_modules detected!")
			waitGroup.Add(1)
			go func() {
				scan.Ecosystems = append(scan.Ecosystems, npm.NewProjectFromNodeModules(scan.Path))
				waitGroup.Done()
			}()
		}
	}

	if golang.DetectGoMod(scan.Path) {
		if golang.DetectGoSum(scan.Path) {
			scan.Log("go.sum detected!")
			waitGroup.Add(1)
			go func() {
				scan.Ecosystems = append(scan.Ecosystems, golang.NewProjectFromGoSum(scan.Path))
				waitGroup.Done()
			}()
		}
	}

	waitGroup.Wait()
}

func (scan *Scan) Log(msg string) {
	if scan.Verbose {
		fmt.Println(msg)
	}
}
