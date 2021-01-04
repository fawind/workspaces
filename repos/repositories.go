package repos

import (
	"context"
	"github.com/fawind/workspaces/config"
	"github.com/google/go-github/v33/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

func GetRepositories(baseUrl url.URL, org string) ([]*github.Repository, error) {
	secrets, err := config.ReadSecrets()
	if err != nil {
		return nil, err
	}

	client := newClient(baseUrl, getTokenForOrg(secrets, baseUrl))

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

func newClient(baseUrl url.URL, token string) *github.Client {
	var c *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		c = oauth2.NewClient(context.Background(), ts)
	}

	client := github.NewClient(c)
	client.BaseURL = &baseUrl
	return client
}

func getTokenForOrg(secrets []config.Secret, baseUrl url.URL) string {
	for _, secret := range secrets {
		if secret.Endpoint.GetApiUrl().String() == baseUrl.String() {
			return secret.Token
		}
	}
	return ""
}
