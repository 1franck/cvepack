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
		fmt.Println(config.Default.NameAndVersion())

		if IsDatabaseUpdateAvailable() {
			UpdateDatabase()
		}

		fmt.Printf("Scanning %s ..\n", args[0])
		if err := common.ValidateDirectory(args[0]); err != nil {
			fmt.Printf("path %s not found\n", args[0])
			return
		}

		scanJob := scan.NewScan(args[0])
		scanJob.Verbose = true
		scanJob.Run()

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

		for _, ecosystem := range scanJob.Ecosystems {
			fmt.Printf(" [%s] %d package(s) analyzed, ", ecosystem.Name(), len(ecosystem.Packages()))
			pkgsVul := scan.Results{}
			for _, pkg := range ecosystem.Packages() {
				vulnerabilities, err := pkgVulQuerier.Query(ecosystem.Name(), pkg.Name(), pkg.Version())
				if err != nil {
					log.Fatal(err)
				}

				if vulnerabilities.IsEmpty() {
					continue
				}

				pkgsVul.Append(pkg, vulnerabilities)
			}

			problemsWord := "problem"
			packageAffectedCount := pkgsVul.UniqueResultCount()
			if packageAffectedCount == 0 {
				fmt.Printf("no %s found\n", problemsWord)
				continue
			}

			if packageAffectedCount > 1 {
				problemsWord += "s"
			}

			fmt.Printf("%d %s found:\n", packageAffectedCount, problemsWord)

			printedDep := make(map[string]bool)
			longestPackageName := pkgsVul.LongestPackageName() + 5
			for _, result := range pkgsVul {
				if _, ok := printedDep[result.Query.ToString()]; !ok {
					fmt.Printf("  [%s%s%s] %s %s %s %s\n",
						packageColor.Sprint(result.Query.Name),
						infoColor.Sprintf("@"),
						versionColor.Sprintf(result.Query.Version),
						strings.Repeat(".", longestPackageName-result.Query.StringLen()),
						colorizeSeveritySummary(result.Vulnerabilities),
						strings.Repeat(".", 35-len(result.Vulnerabilities.SeveritiesSummary())),
						infoColor.Sprint(result.Vulnerabilities.AliasesSummary()))
					printedDep[result.Query.ToString()] = true
				}
			}
		}
	},
}

func colorizeSeveritySummary(vul search.PackageVulnerabilities) string {
	var severitySummary []string
	for level, count := range vul.SeveritiesSummaryMap() {
		level = strings.ToLower(level)
		switch level {
		case "critical":
			severitySummary = append(severitySummary, severityCriticalLeverColor.Sprintf("%d %s", count, level))
			break
		case "high":
			severitySummary = append(severitySummary, severityHighLeverColor.Sprintf("%d %s", count, level))
			break
		default:
			severitySummary = append(severitySummary, severityModerateLeverColor.Sprintf("%d %s", count, level))
			break
		}
	}
	return strings.Join(severitySummary, ", ")
}
