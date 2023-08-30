package cmd

import (
	"cvepack/core/common"
	"cvepack/core/config"
	"cvepack/core/scan"
	"cvepack/core/search"
	"cvepack/core/sqlite"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var showDetails bool

var ScanCommand = &cobra.Command{
	Use:   "scan",
	Short: "Scan a folder",
	Long:  "Scan a folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Default.NameAndVersion())

		if IsDatabaseUpdateAvailable() {
			UpdateDatabase()
		}

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

		for _, path := range args {
			fmt.Printf("Scanning %s ..\n", path)
			if err := common.ValidateDirectory(path); err != nil {
				fmt.Printf("path %s not found\n", path)
				return
			}
			scanPath(path, db)
		}
	},
}

func init() {
	ScanCommand.Flags().BoolVarP(&showDetails, "details", "d", false, "show details")
}

func scanPath(path string, db *sql.DB) {
	scanJob := scan.NewScan(path)
	scanJob.Verbose = true
	scanJob.Run()

	pkgVulQuerier := search.PackageVulnerabilityQuerier(db)

	for _, project := range scanJob.Projects {
		fmt.Printf(" [%s] %d package(s) analyzed, ", project.Ecosystem(), len(project.Packages()))
		pkgsVul := scan.Results{}
		for _, pkg := range project.Packages() {
			vulnerabilities, err := pkgVulQuerier.Query(project.Ecosystem(), pkg.Name(), pkg.Version())
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

				if showDetails {
					for _, vul := range result.Vulnerabilities {
						fmt.Printf("    (%s) %s\n", vul.VulnerabilityId, vul.Summary)
					}
					fmt.Println()
				}
			}
		}

		fmt.Println()
	}
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
