package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Manage the host list",
	Long: `Manages the hosts lists for pScan
	
Add hosts with add command
Delete hosts with the delete command
List hosts with the list command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hosts called")
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)
}
