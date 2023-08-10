package cmd

import (
	"fmt"
	"github.com/1franck/cvepack/internal/config"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Current %s version", config.Default.Name),
	Long:  fmt.Sprintf("Current %s version", config.Default.Name),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s\n", config.Default.Name, config.Default.Version)
	},
}
