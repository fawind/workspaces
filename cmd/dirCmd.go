package cmd

import (
	"fmt"
	"github.com/fawind/workspaces/finder"
	"github.com/fawind/workspaces/shell"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DirCmd)
}

var DirCmd = &cobra.Command{
	Use:  "dir",
	Short: "Print the local directory of the repository",
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
