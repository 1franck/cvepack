package scan

import (
	"cvepack/core/ecosystem"
	"cvepack/core/ecosystem/cratesio"
	"cvepack/core/ecosystem/golang"
	"cvepack/core/ecosystem/npm"
	"cvepack/core/ecosystem/nuget"
	"cvepack/core/ecosystem/packagist"
	"cvepack/core/ecosystem/pypi"
	"cvepack/core/ecosystem/rubygems"
	"fmt"
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

func (scan *Scan) Log(msg string) {
	if scan.Verbose {
		fmt.Println(msg)
	}
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

	if rubygems.DetectGemFile(scan.Path) {
		if rubygems.DetectGemFileLock(scan.Path) {
			scan.Log("Gemfile.lock detected!")
			waitGroup.Add(1)
			go func() {
				scan.Projects = append(scan.Projects, rubygems.NewProjectFromGemFileLock(scan.Path))
				waitGroup.Done()
			}()
		}
	}

	if pypi.DetectPyProjectToml(scan.Path) {
		if pypi.DetectPoetryLock(scan.Path) {
			scan.Log("poetry.lock detected!")
			waitGroup.Add(1)
			go func() {
				scan.Projects = append(scan.Projects, pypi.NewProjectFromPoetryLock(scan.Path))
				waitGroup.Done()
			}()
		} else if pypi.DetectPdmLock(scan.Path) {
			scan.Log("pdm.lock detected!")
			waitGroup.Add(1)
			go func() {
				scan.Projects = append(scan.Projects, pypi.NewProjectFromPdmLock(scan.Path))
				waitGroup.Done()
			}()
		}
	}

	slnFile := nuget.DetectSln(scan.Path)
	if slnFile != "" {
		scan.Log("Sln file detected!")
		waitGroup.Add(1)
		go func() {
			scan.Projects = append(scan.Projects, nuget.NewProjectFromSln(scan.Path, slnFile))
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
