package testers

import (
	"context"
	"github.com/catena-x/gh-org-checks/pkg/data"
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

func (checker HelmChartTester) PerformTest(repoName string) data.RepositoryReport {
	log.Infof("perform helmchart test on repo %s", repoName)
	return checker.contentTest(repoName)
}
