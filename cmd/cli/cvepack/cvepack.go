package main

import (
	"cvepack/core/cmd"
	"fmt"
	"os"
)

func main() {
	//startedAt := time.Now()
	cmd.RootCmd.AddCommand(cmd.VersionCmd)
	cmd.RootCmd.AddCommand(cmd.UpdateCmd)
	cmd.RootCmd.AddCommand(cmd.ScanCommand)
	cmd.RootCmd.AddCommand(cmd.SearchCmd)
	cmd.RootCmd.AddCommand(cmd.StatsCmd)

	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Printf("Finished in %s\n", time.Since(startedAt))
}
