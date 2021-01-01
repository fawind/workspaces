package finder

import (
	"fmt"
	"github.com/fawind/workspaces/workspaces/config"
	"github.com/pkg/errors"
	"path"
)
import "github.com/ktr0731/go-fuzzyfinder"

type searchItem struct {
	organization config.Organization
	repoName     string
}

type SelectedRepo struct {
	Organization config.Organization
	RepoName     string
	LocalDir     string
}

func (s searchItem) String() string {
	return fmt.Sprintf("%s/%s", s.organization, s.repoName)
}

func GetRepository() (SelectedRepo, error) {
	repoCache, err := config.ReadRepoCache()
	if err != nil {
		return SelectedRepo{}, err
	}

	var searchItems []searchItem
	for org, repos := range repoCache {
		for _, repo := range repos {
			searchItems = append(searchItems, searchItem{org, repo})
		}
	}

	idx, err := fuzzyfinder.Find(
		searchItems,
		func(i int) string {
			return searchItems[i].String()
		})
	if err != nil {
		return SelectedRepo{}, err
	}

	selected := searchItems[idx]

	localDir, err := getLocalDirForRepo(selected)
	if err != nil {
		return SelectedRepo{}, err
	}

	return SelectedRepo{selected.organization, selected.repoName, localDir}, nil
}

func getLocalDirForRepo(item searchItem) (string, error) {
	conf, err := config.ReadConfig()
	if err != nil {
		return "", err
	}

	for _, ws := range conf.Workspaces {
		if ws.Organization == item.organization {
			return path.Join(ws.Directory, item.repoName), nil
		}
	}
	return "", errors.Errorf("Could not find local dir for %s in config file", item.organization)
}
