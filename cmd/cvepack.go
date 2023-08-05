package main

import (
	"fmt"
	"github.com/1franck/cvepack/internal/cmd"
	"os"
)

func main() {
	cmd.RootCmd.AddCommand(cmd.VersionCmd)

	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
