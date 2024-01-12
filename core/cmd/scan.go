package cmd

import (
	"cvepack/core/analysis"
	"cvepack/core/common"
	"cvepack/core/config"
	"cvepack/core/database"
	es "cvepack/core/ecosystem"
	"cvepack/core/scan"
	"cvepack/core/search"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	showDetails bool
	url         bool
	silent      bool
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

		db, closeDb := database.ConnectToDefault()
		defer closeDb(db)

		for _, path := range args {
			fmt.Printf("Scanning %s ...\n", path)
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
	ScanCommand.Flags().BoolVarP(&url, "url", "u", false, "url instead of path")
	ScanCommand.Flags().BoolVarP(&silent, "silent", "s", false, "silent")
}

func scanPath(path string, db *sql.DB) {
	verbose := true
	if silent {
		verbose = false
	}

	sourceType := es.PathSource
	if url {
		sourceType = es.UrlSource
	}

	source := es.NewSource(path, sourceType)
	scanResults := scan.Inspect(source)

	_printf := func(format string, a ...interface{}) {
		if verbose {
			fmt.Printf(format, a...)
		}
	}

	_println := func(a ...interface{}) {
		if verbose {
			fmt.Println(a...)
		}
	}

	pkgVulQuerier := search.PackageVulnerabilityQuerier(db)

	for _, project := range scanResults.Projects {
		_printf(
			" [%s] %d %s found, ",
			project.Ecosystem(),
			len(project.Packages()),
			common.Plural(len(project.Packages()), "package", "packages"))

		pkgsVul := scan.PackagesVulnerabilitiesResult{}
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

		packageAffectedCount := pkgsVul.UniqueResultCount()
		if packageAffectedCount == 0 {
			_println("no problem found")
			continue
		}

		_printf(
			"%d %s detected:\n",
			packageAffectedCount,
			common.Plural(packageAffectedCount, "problem", "problems"))

		printedDep := make(map[string]bool)
		longestPackageName := pkgsVul.LongestPackageName() + 5

		for _, result := range pkgsVul {
			if _, ok := printedDep[result.Query.ToString()]; !ok {

				versionToUpdate := analysis.VersionToUpdate(result)
				fmt.Printf("Version to update: %s\n", versionToUpdate)

				aliases := infoColor.Sprint(result.Vulnerabilities.AliasesSummary())
				if showDetails {
					aliases = ""
				}

				_printf("  [%s%s%s] %s %s %s %s\n",
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

						hasFix, fixedVersion := vul.HasFix()
						hasFixText := "no fix"
						if hasFix {
							hasFixText = "fixed in v" + fixedVersion
						}
						_printf("    | %s| %s [%s]\n", alias, vul.Summary, hasFixText)
					}
					_println()
				}
			}
		}

		_println()
		_printf("duration: %s\n", scanResults.Duration())
	}
}

func colorizeSeveritySummary(vul search.PackageVulnerabilities) string {
	var severitySummary []string
	for level, count := range vul.SeveritiesSummaryMap() {
		level = strings.ToLower(level)
		switch level {
		case "critical":
			severitySummary = append(severitySummary, severityCriticalLeverColor.Sprintf("%d %s", count, level))
		case "high":
			severitySummary = append(severitySummary, severityHighLeverColor.Sprintf("%d %s", count, level))
		default:
			severitySummary = append(severitySummary, severityModerateLeverColor.Sprintf("%d %s", count, level))
		}
	}
	return strings.Join(severitySummary, ", ")
}
