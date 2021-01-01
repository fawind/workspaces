package repos

import (
	"context"
	"github.com/google/go-github/v33/github"
	"github.com/pkg/errors"
	"net/url"
)

func GetRepositories(baseUrl url.URL, org string) ([]*github.Repository, error) {
	client := newClient(baseUrl)

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(context.Background(), org, opt)
		if err != nil {
			return nil, errors.Wrapf(err, "error listing repos for org %s/%s", baseUrl, org)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allRepos, nil
}

func newClient(baseUrl url.URL) *github.Client {
	client := github.NewClient(nil)
	client.BaseURL = &baseUrl
	return client
}
