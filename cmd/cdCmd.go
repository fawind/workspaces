package cmd

import (
	"github.com/fawind/workspaces/workspaces/finder"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CdCmd)
}

var CdCmd = &cobra.Command{
	Use:  "cd",
	RunE: changeDir,
}

func changeDir(cmd *cobra.Command, args []string) error {
	println("Changing Directory...")
	return finder.GetRepository()
}
