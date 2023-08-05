package cmd

import (
	"fmt"
	"github.com/1franck/cvepack/internal"
	"github.com/spf13/cobra"
	"strings"
)

var rooCmdLongDescSlice = []string{
	fmt.Sprintf("%s is a tool to detect CVEs in packages.", internal.NAME),
	"It use GitHub Advisory Database to search for CVEs.",
}

var RootCmd = &cobra.Command{
	Use:   strings.ToLower(internal.NAME_VERSION),
	Short: fmt.Sprintf("%s is a tool to detect CVEs in packages", internal.NAME),
	Long:  strings.Join(rooCmdLongDescSlice, " "),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootCmdCliHeader())
		fmt.Println("Available Commands:")
		for _, subCmd := range cmd.Commands() {
			if subCmd.Name() == "help" || subCmd.Name() == "completion" {
				continue
			}
			fmt.Printf("  %-10s %s\n", subCmd.Name(), subCmd.Short)
		}
	},
}

func rootCmdCliHeader() string {
	maxLength := 0
	for _, line := range rooCmdLongDescSlice {
		if len(strings.TrimSpace(line)) > maxLength {
			maxLength = len(line) + 1
		}
	}
	var header []string
	header = append(header, strings.Repeat("=", maxLength))
	for _, line := range rooCmdLongDescSlice {
		if line != "" {
			header = append(header, strings.TrimSpace(line)+".")
		}
	}
	header = append(header, strings.Repeat("=", maxLength))
	return strings.Join(header, "\n")
}
