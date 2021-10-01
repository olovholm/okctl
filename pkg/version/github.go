package version

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
)

type githubb struct {
	Client *github.Client
}

type repositoryRelease = github.RepositoryRelease

const listReleasesPageSize = 100

// ListReleases lists the given repository's releases
func (g *githubb) ListReleases(ctx context.Context, owner, repo string) ([]*repositoryRelease, error) {
	opts := &github.ListOptions{
		PerPage: listReleasesPageSize,
	}

	var allReleases []*repositoryRelease

	for {
		// Documentation: https://docs.github.com/en/rest/reference/repos#list-release-assets
		releases, response, err := g.Client.Repositories.ListReleases(ctx, owner, repo, opts)
		if err != nil {
			return nil, fmt.Errorf("listing releases: %w", err)
		}

		allReleases = append(allReleases, releases...)

		if response.NextPage == 0 {
			break
		}

		opts.Page = response.NextPage
	}

	return allReleases, nil
}

func newGithub() *githubb {
	return &githubb{
		Client: github.NewClient(nil),
	}
}
