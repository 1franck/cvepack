package cmd

import (
	"cvepack/core/cli"
	"cvepack/core/config"
	"cvepack/core/database"
	"cvepack/core/stats"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Vulnerabilities database stats",
	Long:  "Vulnerabilities database stats",
	Run: func(cmd *cobra.Command, args []string) {
		cli.PrintNameWithVersionHeader()

		db, closeDb := database.ConnectToDefault()
		defer closeDb(db)

		total := stats.GetTotalVulnerabilities(db)
		cli.PrintWithUnderline(fmt.Sprintf("Total vulnerabilities: %d", total))

		ecosystemStats, err := stats.GetEcosystemsStats(db)
		if err != nil {
			log.Fatalf("error while getting ecosystem stats: %s", err)
		}

		ecosystemStatsFormatted := make(map[string]string, len(ecosystemStats))
		for ecosystem, count := range ecosystemStats {
			language := stats.Ecosystems[ecosystem]
			if language != ecosystem {
				language = fmt.Sprintf("%s (%s)", ecosystem, language)
			}
			ecosystemStatsFormatted[language] = fmt.Sprintf("%d", count)

			percentageRounded := fmt.Sprintf("%.1f", float64(count)/float64(total)*100)
			ecosystemStatsFormatted[language] = ecosystemStatsFormatted[language] + " (" + percentageRounded + "%)"
		}

		cli.PrintMap(ecosystemStatsFormatted, nil)

		lastModified, err := database.LastModified(config.Default.DatabaseFilePath())
		if err != nil {
			log.Fatalf("error while getting database last modified time: %s", err)
		}

		cli.PrintWithUpperLine(fmt.Sprintf("Last update: %s", lastModified))
	},
}
