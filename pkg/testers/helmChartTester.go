package testers

import (
	"context"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type HelmChartTester struct {
	TestProperty
	ContentTester
}

func NewHelmChartTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new helmchart tester")
	return HelmChartTester{
		TestProperty: NewTestProperty("Helm Chart"),
		ContentTester: NewContentTester(ctx, owner, githubClient,
			[]repositoryContent{
				{
					path:        "charts",
					contentType: Directory,
				},
			}),
	}
}
