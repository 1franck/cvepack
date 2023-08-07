// Take a compiled db and upload it to github.com/1franck/cvepack-database repository
package main

import (
	"database/sql"
	"flag"
	"github.com/1franck/cvepack/internal/common"
	"github.com/1franck/cvepack/internal/common/checksum"
	"github.com/1franck/cvepack/internal/core"
	"github.com/1franck/cvepack/internal/git"
	"github.com/1franck/cvepack/internal/sqlite"
	"log"
	"os"
	"path/filepath"
)

func main() {
	start := common.TimerStart()

	advDbFilePath := flag.String("src", "", "Filepath of advisory database")
	compiledAdvRepoPath := flag.String("dst", "", "Path of compiled advisory database repository")

	flag.Parse()

	if *advDbFilePath == "" || *compiledAdvRepoPath == "" {
		flag.Usage()
		return
	}

	if !common.FileExists(*advDbFilePath) {
		log.Fatalf("Advisory database not found at %s", *advDbFilePath)
	}

	// Check database if ok
	db, err := sqlite.Connect(*advDbFilePath)
	defer func(db *sql.DB) {
		err := db.Close()
		log.Fatal(err)
	}(db)
	if err != nil {
		log.Fatal(err)
	}

	vulCount, err := core.CountVulnerabilities(db)
	if err != nil {
		log.Fatal(err)
	}
	if vulCount == 0 {
		log.Fatal("No vulnerabilities found in database, Abort!")
	}
	log.Printf("Found %d vulnerabilities in database...\n", vulCount)

	// Calculate checksum of advisory database and write it to db.checksum
	cs, err := checksum.FromFile(*advDbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Checksum of %s: %s\n", *advDbFilePath, cs)

	checksumFile := filepath.Join(*compiledAdvRepoPath, "db.checksum")
	err = os.WriteFile(checksumFile, []byte(cs), 0644)
	if err != nil {
		log.Fatalf("Error writing checksum to %s: %s", checksumFile, err)
	}

	// Copy db to repo
	destinationFile, err := os.Create(filepath.Join(*compiledAdvRepoPath, "advisories.db"))
	if err != nil {
		log.Fatalf("Error creating destination file: %s", err)
	}
	defer destinationFile.Close()

	// Stage all modified files
	result, err := git.StageAllModified(*compiledAdvRepoPath)
	if err != nil {
		log.Fatalf("Error staging files: %s", err)
	}
	log.Println(result)

	// Commit with a message
	commitMsg := "Auto-commit: Update database"
	result, err = git.Commit(*compiledAdvRepoPath, commitMsg)
	if err != nil {
		log.Fatalf("Error committing: %s\n%s", err, result)
	}
	log.Println(result)

	// Push to origin/main
	result, err = git.Push(*compiledAdvRepoPath, "origin", "main")
	if err != nil {
		log.Fatalf("Error pushing: %s", err)
	}
	log.Println(result)

	log.Println("Done!")
	common.PrintTimer(start)
}