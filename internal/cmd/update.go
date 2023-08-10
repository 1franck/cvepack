package cmd

import (
	"fmt"
	"github.com/1franck/cvepack/internal/config"
	"github.com/1franck/cvepack/internal/core"
	"github.com/1franck/cvepack/internal/update"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update CVE database",
	Long:  "Update CVE database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating CVE database ...")
		update.UpdateDatabase("./")
		fmt.Println("Checking ...")
		err := core.IsDatabaseOk(config.Default.DatabaseFilePath())
		if err != nil {
			_ = fmt.Errorf("error checking database: %s", err)
		}
		fmt.Println("... Database OK!")
	},
}
