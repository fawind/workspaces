package cmd

import (
	"fmt"
	"github.com/fawind/workspaces/workspaces/finder"
	"github.com/fawind/workspaces/workspaces/shell"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DirCmd)
}

var DirCmd = &cobra.Command{
	Use:  "dir",
	RunE: printLocalDir,
}

func printLocalDir(cmd *cobra.Command, args []string) error {
	println("Changing Directory...")
	repo, err := finder.GetRepository()
	if err != nil {
		return err
	}
	if err := shell.MayCloneRepo(repo); err != nil {
		return err
	}

	fmt.Println(repo.LocalDir)
	return nil
}
