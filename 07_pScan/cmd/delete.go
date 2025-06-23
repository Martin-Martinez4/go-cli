/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/Martin-Martinez4/go-cli/pScan/scan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func deleteAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, h := range args {
		if err := hl.Remove(h); err != nil {
			return err
		}
		fmt.Fprintln(out, "Deleted host:", h)
	}
	return hl.Save(hostsFile)
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:          "delete",
	Aliases:      []string{"d"},
	SilenceUsage: true,
	Args:         cobra.MinimumNArgs(1),
	Short:        "Delete host(s) from the list",
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile := viper.GetString("hosts-file")

		return deleteAction(os.Stdout, hostsFile, args)
	},
}

func init() {
	hostCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
