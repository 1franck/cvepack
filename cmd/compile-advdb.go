package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"github.com/1franck/cvepack/internal/common"
	"github.com/1franck/cvepack/internal/core"
	"github.com/1franck/cvepack/internal/git"
	"github.com/1franck/cvepack/internal/osv"
	"github.com/1franck/cvepack/internal/sqlite"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

func main() {
	start := common.TimerStart()

	if !git.CommandExists() {
		log.Fatal("Git command not found, please install it")
	}

	advDbPath := flag.String("src", "", "Path of advisory database repository")
	onlyReviewedFlag := flag.Bool("only-reviewed", true, "Only scan reviewed advisories")
	outputDbFlag := flag.String("output", "", "Output database file")

	flag.Parse()

	if *advDbPath == "" {
		flag.Usage()
		return
	}

	if *outputDbFlag == "" {
		*outputDbFlag = "./advisories.db"
	} else if *outputDbFlag == "./" || *outputDbFlag == "." || *outputDbFlag != "/" {
		log.Fatal("Output database file is not valid, should be a file")
	}

	err := common.ValidateDirectory(*advDbPath)
	logFatal(err)

	log.Printf("Pulling latest changes from %s ...", *advDbPath)
	pullResult, err := git.Pull(*advDbPath)
	logFatal(err)

	log.Printf(pullResult)

	pathToScan := *advDbPath
	if *onlyReviewedFlag {
		pathToScan = filepath.Join(*advDbPath, "advisories/github-reviewed")
	}

	log.Printf("Scanning %s ...", pathToScan)

	jsonFiles := []string{}
	err = filepath.Walk(pathToScan,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".json" {
				jsonFiles = append(jsonFiles, path)
			}
			return nil
		})
	logFatal(err)

	log.Printf("Found %d JSON files", len(jsonFiles))
	log.Printf("Preparing %s database ...", *outputDbFlag)

	if common.FileExists(*outputDbFlag) {
		err = os.Remove(*outputDbFlag)
		if err != nil {
			log.Fatal(err)
		}
	}

	db, err := sqlite.Connect(*outputDbFlag)
	defer func(db *sql.DB) {
		err := db.Close()
		logFatal(err)
	}(db)
	logFatal(err)

	_, err = db.Exec(core.DbSchema)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Processing JSON files to database ...")

	for _, jsonFile := range jsonFiles {
		jsonFileContent, err := os.ReadFile(jsonFile)
		if err != nil {
			log.Println(err)
			continue
		}
		vul := osv.Osv{}
		err = json.Unmarshal(jsonFileContent, &vul)
		if err != nil {
			log.Printf("Error decoding %s : %s\n", jsonFile, err)
			continue
		}
		err = core.InsertVulnerability(db, vul)
		if err != nil {
			log.Printf("Error inserting %s : %s\n", jsonFile, err)
			continue
		}
	}
	log.Println("Done!")

	common.PrintTimer(start)
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
