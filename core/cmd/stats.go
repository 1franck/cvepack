package cmd

import (
	"cvepack/core/cli"
	"cvepack/core/common"
	"cvepack/core/config"
	"cvepack/core/database"
	"cvepack/core/stats"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
	"log"
)

var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Packages vulnerabilities stats",
	Long:  "Packages vulnerabilities stats",
	Run: func(cmd *cobra.Command, args []string) {

		db, closeDb := database.ConnectToDefault()
		defer closeDb(db)

		cli.PrintNameWithVersionHeader()

		total := stats.GetTotalVulnerabilities(db)
		fmt.Printf("Total vulnerabilities: %d\n", total)

		lastModified, err := database.LastModified(config.Default.DatabaseFilePath())
		if err != nil {
			log.Fatalf("error while getting database last modified time: %s", err)
		}
		fmt.Printf("Last update: %s\n", lastModified)

		ecosystemStats, err := stats.GetEcosystemsStats(db)
		if err != nil {
			log.Fatalf("error while getting ecosystem stats: %s", err)
		}

		rows := make([][]string, 0)

		for _, es := range ecosystemStats {
			rows = append(rows, []string{
				common.DefaultTextPad(es.GetTitle()),
				common.DefaultTextPad(fmt.Sprintf("%5d", es.Count)),
				common.DefaultTextPad(fmt.Sprintf("%5.2f %s", es.Percentage(total), "%")),
			})
		}

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#fff"))).
			Headers(
				common.DefaultTextPad("Ecosystem"),
				common.DefaultTextPad("Count"),
				common.DefaultTextPad("DB ratio")).
			Rows(rows...)

		fmt.Println(t.Render())
	},
}
