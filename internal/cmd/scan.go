package cmd

import (
	"database/sql"
	"fmt"
	"github.com/1franck/cvepack/internal/common"
	"github.com/1franck/cvepack/internal/config"
	"github.com/1franck/cvepack/internal/core/search"
	"github.com/1franck/cvepack/internal/scan"
	"github.com/1franck/cvepack/internal/sqlite"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var ScanCommand = &cobra.Command{
	Use:   "scan",
	Short: "Scan a folder",
	Long:  "Scan a folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("scanning %s\n", args[0])
		if err := common.ValidateDirectory(args[0]); err != nil {
			fmt.Printf("path %s not found\n", args[0])
			return
		}

		scan := scan.NewScan(args[0])
		scan.Verbose = true
		scan.Run()

		db, err := sqlite.Connect(config.Default.DatabaseFilePath())
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Printf("error while closing database: %s", err)
				log.Fatal(err)
			}
		}(db)
		if err != nil {
			log.Printf("error while connecting to database: %s", err)
			log.Fatal(err)
		}

		pkgVulQuerier := search.PackageVulnerabilityQuerier(db)

		for _, ecosystem := range scan.Ecosystems {
			fmt.Printf("[%s] %d package(s) analyzed, ", ecosystem.Name(), len(ecosystem.Packages()))
			pkgsVul := make(ecosytemVulnerabilities)
			for _, pkg := range ecosystem.Packages() {
				vulnerabilities, err := pkgVulQuerier.Query(ecosystem.Name(), pkg.Name(), pkg.Version())
				if err != nil {
					log.Fatal(err)
				}

				if vulnerabilities.IsEmpty() {
					continue
				}

				pkgsVul[pkg.Name()] = vulnerabilities
			}

			if len(pkgsVul) > 0 {
				if len(pkgsVul) == 1 {
					fmt.Printf("%d problem found:\n", len(pkgsVul))
				} else {
					fmt.Printf("%d problems found:\n", len(pkgsVul))
				}
			} else {
				fmt.Printf("no problem found\n")
			}

			longuestPackageName := pkgsVul.LongestPackageName() + 10
			for pkgName, vulnerabilities := range pkgsVul {
				fmt.Printf(" ├ %s %s %s %s %s\n",
					pkgName,
					strings.Repeat(".", longuestPackageName-len(pkgName)),
					vulnerabilities.SeveritiesSummary(),
					strings.Repeat(".", 30-len(vulnerabilities.SeveritiesSummary())),
					vulnerabilities.AliasesSummary())
			}

		}
	},
}

type ecosytemVulnerabilities map[string]search.PackageVulnerabilities

func (e ecosytemVulnerabilities) LongestPackageName() int {
	var longest int
	for k := range e {
		if len(k) > longest {
			longest = len(k)
		}
	}
	return longest
}
