package finder

import (
	"fmt"
	"github.com/fawind/workspaces/config"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"path"
)
import "github.com/ktr0731/go-fuzzyfinder"

type searchItem struct {
	organization config.Organization
	repoName     string
}

type LocalRepo struct {
	Organization config.Organization
	RepoName     string
	LocalDir     string
}

func (s searchItem) String() string {
	return fmt.Sprintf("%s/%s", s.organization, s.repoName)
}

func GetRepository() (LocalRepo, error) {
	repoCache, err := config.ReadRepoCache()
	if err != nil {
		return LocalRepo{}, err
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
		return LocalRepo{}, err
	}

	selected := searchItems[idx]

	localDir, err := getLocalDirForRepo(selected)
	if err != nil {
		return LocalRepo{}, err
	}

	return LocalRepo{selected.organization, selected.repoName, localDir}, nil
}

func getLocalDirForRepo(item searchItem) (string, error) {
	conf, err := config.ReadConfig()
	if err != nil {
		return "", err
	}

	for _, ws := range conf.Workspaces {
		if ws.Organization == item.organization {
			path := path.Join(ws.Directory, item.repoName)
			expanded, err := homedir.Expand(path)
			if err != nil {
				return "", errors.Wrapf(err, "error expanding home dir for path: ", path)
			}
			return expanded, nil
		}
	}
	return "", errors.Errorf("Could not find local dir for %s in config file", item.organization)
}
