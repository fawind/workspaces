package config

import (
	"fmt"
	"github.com/fawind/workspaces/workspaces/repos"
	"github.com/google/go-github/v33/github"
)

func RefreshRepoCache() error {
	conf, err := ReadConfig()
	if err != nil {
		return err
	}

	var repoCache = make(RepoCache)
	numRepos := 0
	for _, ws := range conf.Workspaces {
		repos, err := repos.GetRepositories(ws.Organization.GetApiUrl(), ws.Organization.GetOrgName())
		if err != nil {
			return err
		}
		repoCache[ws.Organization] = getRepoNames(repos)
		numRepos += len(repos)
	}

	err = WriteRepoCache(repoCache)

	println("Updated the local repo cache:")
	for org, repos := range repoCache {
		fmt.Printf("%s: %d repos\n", org, len(repos))
	}

	return err
}

func getRepoNames(repos []*github.Repository) []string {
	var repoNames []string
	for _, repo := range repos {
		repoNames = append(repoNames, *repo.Name)
	}
	return repoNames
}
