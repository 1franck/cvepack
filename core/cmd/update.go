package cmd

import (
	"cvepack/core/config"
	"cvepack/core/database"
	"cvepack/core/update"
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update CVE database",
	Long:  "Update CVE database",
	Run: func(cmd *cobra.Command, args []string) {
		UpdateDatabase()
	},
}

var databaseUpdateAvailable = false
var databaseUpdateAvailableChecked = false

func IsDatabaseUpdateAvailable() bool {
	if databaseUpdateAvailableChecked {
		return databaseUpdateAvailable
	}
	if updateAvailable, _ := update.IsAvailable(config.Default); updateAvailable {
		databaseUpdateAvailable = true
		databaseUpdateAvailableChecked = true
	}
	return databaseUpdateAvailable
}

func UpdateDatabase() {
	fmt.Println("Updating CVE database ...")
	update.UpdateDatabase("./")
	fmt.Println("Checking ...")
	err := database.IsDatabaseOk(config.Default.DatabaseFilePath())
	if err != nil {
		_ = fmt.Errorf("error checking database: %s", err)
	}
	fmt.Println("... Database OK!")
}
