package finder

import (
	"fmt"
	"github.com/fawind/workspaces/workspaces/config"
)
import "github.com/ktr0731/go-fuzzyfinder"

type searchItem struct {
	organization config.Organization
	name         string
}

func (s searchItem) String() string {
	return fmt.Sprintf("%s/%s", s.organization, s.name)
}

func GetRepository() error {
	repoCache, err := config.ReadRepoCache()
	if err != nil {
		return err
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
		return err
	}
	fmt.Printf("selected: %v\n", idx)
	return nil
}
