package cmd

import (
	"fmt"
	"github.com/1franck/cvepack/internal"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Current %s version", internal.NAME),
	Long:  fmt.Sprintf("Current %s version", internal.NAME),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s\n", internal.NAME, internal.VERSION)
	},
}
