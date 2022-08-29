package testers

import (
	"context"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type HelmChartTester struct {
	ContentTester
}

func NewHelmChartTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new helmchart tester")
	return HelmChartTester{ContentTester{
		testType:     "HelmChartCheck",
		ctx:          ctx,
		owner:        owner,
		githubClient: githubClient,
		contents: []repositoryContent{
			{
				path:        "charts",
				contentType: Directory,
			},
		},
	}}
}
