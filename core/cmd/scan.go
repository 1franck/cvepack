package cmd

import (
	"cvepack/core/common"
	"cvepack/core/config"
	"cvepack/core/database"
	"cvepack/core/scan"
	"cvepack/core/search"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"time"
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

		db, closeDb := database.Connect()
		defer closeDb(db)

		for _, path := range args {
			fmt.Printf("Scanning %s ... at %s\n", path, time.Now().Format("2006-01-02 15:04:05"))
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
		fmt.Printf(" [%s] %d package(s) found, ", project.Ecosystem(), len(project.Packages()))
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

		fmt.Printf("%d %s detected:\n", packageAffectedCount, problemsWord)

		printedDep := make(map[string]bool)
		longestPackageName := pkgsVul.LongestPackageName() + 5
		for _, result := range pkgsVul {
			if _, ok := printedDep[result.Query.ToString()]; !ok {

				aliases := infoColor.Sprint(result.Vulnerabilities.AliasesSummary())
				if showDetails {
					aliases = ""
				}

				fmt.Printf("  [%s%s%s] %s %s %s %s\n",
					packageColor.Sprint(result.Query.Name),
					infoColor.Sprintf("@"),
					versionColor.Sprintf(result.Query.Version),
					strings.Repeat(".", longestPackageName-result.Query.StringLen()),
					colorizeSeveritySummary(result.Vulnerabilities),
					strings.Repeat(".", 35-len(result.Vulnerabilities.SeveritiesSummary())),
					aliases)

				printedDep[result.Query.ToString()] = true

				if showDetails {
					for _, vul := range result.Vulnerabilities {
						alias := vul.VulnerabilityId
						if len(vul.AliasesParsed()) > 0 {
							alias = vul.AliasesToString()
						}
						if len(alias) < 20 {
							alias += strings.Repeat(" ", 20-len(alias))
						}

						hasFix := "no fix"
						if vul.VersionFixed != nil && strings.TrimSpace(*vul.VersionFixed) != "" {
							hasFix = "fixed in v" + *vul.VersionFixed
						}
						fmt.Printf("    | %s| %s [%s]\n", alias, vul.Summary, hasFix)
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
