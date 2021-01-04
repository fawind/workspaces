package cmd

import (
	"fmt"
	"github.com/fawind/workspaces/workspaces/finder"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(GithubLinkCmd)
}

var GithubLinkCmd = &cobra.Command{
	Use:  "github",
	RunE: printGithubLink,
}

func printGithubLink(cmd *cobra.Command, args []string) error {
	repo, err := finder.GetRepository()
	if err != nil {
		return err
	}
	fmt.Printf("%s/%s\n", repo.Organization, repo.RepoName)
	return nil
}
