package cmd

import (
	"cvepack/core/config"
	"cvepack/core/search"
	"cvepack/core/sqlite"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a package in database",
	Long:  "Search a package in database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s\n", config.Default.Name, config.Default.Version)

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

		querier := search.LookupPackageQuerier(db)
		results, err := querier.Query(args[0], "")
		if err != nil {
			log.Fatal(err)
		}

		if len(results) == 0 {
			fmt.Printf("No results found for '%s'", args[0])
			return
		}

		fmt.Printf("Found %d results for '%s' ...\n", len(results), args[0])

		for _, result := range results {
			versionFixed := "-"
			if result.VersionFixed != nil {
				versionFixed = *result.VersionFixed
			}

			versionLastAffected := "-"
			if result.VersionLastAffected != nil {
				versionLastAffected = *result.VersionLastAffected
			}

			tag := fmt.Sprintf(
				"[%s - %s]",
				packageColor.Sprint(result.AliasesToString()),
				colorizeSeverityLevel(result.SeverityLevel()),
			)
			if result.AliasesToString() == "" {
				tag = fmt.Sprintf(
					"[%s]",
					colorizeSeverityLevel(result.SeverityLevel()),
				)
			}

			fmt.Printf("\n%s\n Introduced: %s, Fixed: %s, Last Affected: %s\n Ecosystem: %s, Ref: %s\n Summary: %s\n",
				tag,
				result.VersionIntroduced,
				versionFixed,
				versionLastAffected,
				result.PackageEcosystem,
				result.VulnerabilityId,
				result.Summary)
		}
	},
}

func colorizeSeverityLevel(level string) string {
	switch level {
	case "CRITICAL":
		return severityCriticalLeverColor.Sprint(level)
	case "HIGH":
		return severityHighLeverColor.Sprint(level)
	default:
		return severityModerateLeverColor.Sprint(level)
	}
}
