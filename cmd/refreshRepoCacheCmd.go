package cmd

import (
	"github.com/fawind/workspaces/workspaces/config"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(RefreshRepoCacheCmd)
}

var RefreshRepoCacheCmd = &cobra.Command{
	Use:  "refresh",
	RunE: refreshRepoCache,
}

func refreshRepoCache(cmd *cobra.Command, args []string) error {
	println("Refreshing repo cache...")
	return config.RefreshRepoCache()
}
