package cmd

import (
	"fmt"
	"github.com/fawind/workspaces/workspaces/finder"
	"github.com/fawind/workspaces/workspaces/shell"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(BrowserCmd)
}

var BrowserCmd = &cobra.Command{
	Use:  "browser",
	RunE: openInBrowser,
}

func openInBrowser(cmd *cobra.Command, args []string) error {
	repo, err := finder.GetRepository()
	if err != nil {
		return err
	}
	return shell.OpenInBrowser(fmt.Sprintf("%s/%s", repo.Organization, repo.RepoName))
}
