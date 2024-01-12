// CVEPACK - advdb-compiler
// compile latest change of https://github.com/github/advisory-database.git into an sqlite database
package main

import (
	"cvepack/core/common"
	"cvepack/core/database"
	"cvepack/core/git"
	"cvepack/core/osv"
	"cvepack/core/sqlite"
	"database/sql"
	"encoding/json"
	"flag"
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

	advDbPath := flag.String("src", "", "Source of advisory database repository")
	onlyReviewedFlag := flag.Bool("only-reviewed", true, "Only scan reviewed advisories")
	outputDbFlag := flag.String("output", "", "Output database file")

	flag.Parse()

	if *advDbPath == "" {
		log.Println("Please specify the path of advisory database repository (-src)")
		flag.Usage()
		return
	}

	if *outputDbFlag == "" {
		*outputDbFlag = "./advisories.db"
		if !*onlyReviewedFlag {
			*outputDbFlag = "./advisories-unreviewed.db"
		}
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
	} else {
		pathToScan = filepath.Join(*advDbPath, "advisories/unreviewed")
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

	_, err = db.Exec(database.DbSchema)
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
		err = database.InsertVulnerability(db, vul)
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
