package scan

import (
	"fmt"
	"github.com/1franck/cvepack/internal/ecosystem"
	"github.com/1franck/cvepack/internal/ecosystem/cratesio"
	"github.com/1franck/cvepack/internal/ecosystem/golang"
	"github.com/1franck/cvepack/internal/ecosystem/npm"
	"github.com/1franck/cvepack/internal/ecosystem/packagist"
	"sync"
)

type Scan struct {
	Path     string
	Projects []ecosystem.Project
	Verbose  bool
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
				scan.Projects = append(scan.Projects, npm.NewProjectFromPackageLockJson(scan.Path))
				waitGroup.Done()
			}()
		} else if npm.DetectNodeModules(scan.Path) {
			scan.Log("node_modules detected!")
			waitGroup.Add(1)
			go func() {
				scan.Projects = append(scan.Projects, npm.NewProjectFromNodeModules(scan.Path))
				waitGroup.Done()
			}()
		}
	}

	if golang.DetectGoMod(scan.Path) {
		if golang.DetectGoSum(scan.Path) {
			scan.Log("go.sum detected!")
			waitGroup.Add(1)
			go func() {
				scan.Projects = append(scan.Projects, golang.NewProjectFromGoSum(scan.Path))
				waitGroup.Done()
			}()
		}
	}

	if packagist.DetectComposerJson(scan.Path) {
		if packagist.DetectComposerLock(scan.Path) {
			scan.Log("composer.lock detected!")
			waitGroup.Add(1)
			go func() {
				scan.Projects = append(scan.Projects, packagist.NewProjectFromComposerLock(scan.Path))
				waitGroup.Done()
			}()
		}
	}

	if cratesio.DetectCargoToml(scan.Path) {
		if cratesio.DetectCargoLock(scan.Path) {
			scan.Log("Cargo.lock detected!")
			waitGroup.Add(1)
			go func() {
				scan.Projects = append(scan.Projects, cratesio.NewProjectFromCargoLock(scan.Path))
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
