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
	"os"
	"strings"
)

var (
	showDetailsFlag    bool
	urlFlag            bool
	silentFlag         bool
	outputJsonFileFlag string
)

var ScanCommand = &cobra.Command{
	Use:   "scan",
	Short: "Scan a folder",
	Long:  "Scan a folder",
	Run: func(cmd *cobra.Command, args []string) {
		if !silentFlag {
			fmt.Println(config.Default.NameAndVersion())
		}
		if IsDatabaseUpdateAvailable() {
			UpdateDatabase()
		}

		db, closeDb := database.ConnectToDefault()
		defer closeDb(db)

		projectsResult := make(scan.ProjectsVulnerabilitiesResult, 0)
		for _, path := range args {
			scanPath(path, db, &projectsResult)
		}

		if outputJsonFileFlag != "" {
			jsonResult, err := projectsResult.ToJson()
			if err != nil {
				fmt.Println(err)
			} else {
				err = os.WriteFile(outputJsonFileFlag, jsonResult, 0644)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	},
}

func init() {
	ScanCommand.Flags().BoolVarP(&showDetailsFlag, "details", "d", false, "show details")
	ScanCommand.Flags().BoolVarP(&urlFlag, "url", "u", false, "url instead of path")
	ScanCommand.Flags().BoolVarP(&silentFlag, "silent", "s", false, "silent")
	ScanCommand.Flags().StringVarP(&outputJsonFileFlag, "output", "o", "", "output json file")
}

func scanPath(path string, db *sql.DB, projectsResult *scan.ProjectsVulnerabilitiesResult) {

	var (
		err           error
		verbose       = true
		sourceType    = es.PathSource
		pkgVulQuerier = search.PackageVulnerabilityQuerier(db)
	)

	if silentFlag {
		verbose = false
	}

	if urlFlag {
		sourceType = es.UrlSource
	}

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

	source := es.NewSource(path, sourceType)
	if err = es.ValidateSource(source); err != nil {
		_println(err)
		return
	}

	_printf("Scanning %s ...\n", path)

	scanResults := scan.Inspect(source)

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

		projectsResult.Add(project, &pkgsVul)

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
				if showDetailsFlag {
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

				if showDetailsFlag {
					for _, vul := range result.Vulnerabilities {
						alias := vul.VulnerabilityId
						if len(vul.AliasesParsed) > 0 {
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
