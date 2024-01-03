// CVEPACK - advdb-uploader
// Take a compiled sqlite db and upload it to github.com/1franck/cvepack-database repository
package main

import (
	"cvepack/core/common"
	"cvepack/core/common/checksum"
	"cvepack/core/database"
	"cvepack/core/git"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var readmeTpl = `# Database for CVEPack

This repository contains a SQLite database of CVEs compiled from [GitHub Advisory Database](https://github.com/github/advisory-database), which is used by [CVEPack tool](https://github.com/1franck/cvepack).

This repository is updated regularly and programmatically.

Last update: {{ last_update }}`

func main() {
	start := common.TimerStart()

	advDbFilePath := flag.String("src", "", "Filepath of advisory database")
	compiledAdvRepoPath := flag.String("dst", "", "Path of compiled advisory database repository")
	simulation := flag.Bool("simulation", false, "Simulation mode (no gitCommit, no push) [default: false]")

	flag.Parse()

	if *advDbFilePath == "" || *compiledAdvRepoPath == "" {
		flag.Usage()
		return
	}

	if !common.FileExists(*advDbFilePath) {
		log.Fatalf("Advisory database not found at %s", *advDbFilePath)
	}

	var err error

	// Check database if ok
	log.Println("Checking database...")
	err = database.IsDatabaseOk(*advDbFilePath)
	if err != nil {
		log.Fatalf("Error checking database: %s", err)
	}
	log.Println("Database ok!")

	// Calculate checksum of advisory database and write it to db.checksum
	cs, err := checksum.FromFile(*advDbFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Checksum of %s: %s\n", *advDbFilePath, cs)

	// Read old checksum and compare it with the new one
	checksumIsEqual, err := isChecksumEqual(*compiledAdvRepoPath, cs)
	if err != nil {
		log.Fatal(err)
	}

	if checksumIsEqual {
		log.Println("Checksums are equal, no need to update")
		return
	}

	if *simulation {
		log.Println("Simulation mode, no gitCommit, no push")
		return
	}

	// Write the new checksum to db.checksum
	err = writeChecksum(*compiledAdvRepoPath, cs)
	if err != nil {
		log.Fatal(err)
	}

	// Copy db to repo
	err = common.CopyFile(*advDbFilePath, filepath.Join(*compiledAdvRepoPath, "advisories.db"))
	if err != nil {
		log.Fatal("Error copying file:", err)
		return
	}

	// Update README.md
	err = updateReadme(*compiledAdvRepoPath)
	if err != nil {
		log.Fatal(err)
	}

	// Stage all modified files
	stageResult, err := stageAllModified(*compiledAdvRepoPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stageResult)

	// Commit with a message
	commitResult, err := gitCommit(*compiledAdvRepoPath, "Auto-gitCommit: Update database")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(commitResult)

	pushResult, err := gitPush(*compiledAdvRepoPath, "origin", "main")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pushResult)

	log.Println("Done!")
	common.PrintTimer(start)
}

func isChecksumEqual(path string, checksum string) (bool, error) {
	oldChecksumFile, err := os.Open(filepath.Join(path, "db.checksum"))
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error opening db.checksum: %s", err))
	}
	defer func(oldChecksumFile *os.File) {
		err := oldChecksumFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(oldChecksumFile)

	oldChecksum, err := io.ReadAll(oldChecksumFile)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error reading db.checksum: %s", err))
	}

	return string(oldChecksum) == checksum, nil
}

func writeChecksum(path string, checksum string) error {
	checksumFile := filepath.Join(path, "db.checksum")
	err := os.WriteFile(checksumFile, []byte(checksum), 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("Error writing checksum to %s: %s", checksumFile, err))
	}
	return nil
}

func updateReadme(path string) error {
	readmeContent := strings.ReplaceAll(readmeTpl, "{{ last_update }}", time.Now().Format("2006-01-02 15:04:05"))
	err := os.WriteFile(filepath.Join(path, "README.md"), []byte(readmeContent), 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("Error writing README.md: %s", err))
	}
	return nil
}

func stageAllModified(path string) (string, error) {
	result, err := git.StageAllModified(path)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error staging files: %s", err))
	}
	return result, nil
}

func gitCommit(path string, message string) (string, error) {
	result, err := git.Commit(path, message)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error committing: %s\n%s", err, result))
	}
	return result, nil
}

func gitPush(path string, remote string, branch string) (string, error) {
	result, err := git.Push(path, remote, branch)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error pushing: %s", err))
	}
	return result, nil
}
