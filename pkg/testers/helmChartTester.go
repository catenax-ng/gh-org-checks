package testers

import (
	"context"
	"fmt"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type contentType string

const (
	File      contentType = "file"
	Directory contentType = "directory"
)

type repositoryContent struct {
	path    string
	content contentType
}

var (
	contents = []repositoryContent{
		{
			path:    "charts",
			content: Directory,
		},
	}
)

type HelmChartTester struct {
	testType     string
	ctx          context.Context
	owner        string
	githubClient *github.Client
}

func NewHelmChartTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new helmchart tester")
	return HelmChartTester{
		testType:     "HelmChartCheck",
		ctx:          ctx,
		owner:        owner,
		githubClient: githubClient,
	}
}

func (checker HelmChartTester) PerformTest(repoName string) data.RepositoryReport {
	log.Infof("perform helmchart test on repo %s", repoName)

	var testSuccess = true
	var logs = ""

	for _, content := range contents {
		_, _, resp, err := checker.githubClient.Repositories.GetContents(checker.ctx, checker.owner, repoName, content.path, &github.RepositoryContentGetOptions{})

		if resp.StatusCode != http.StatusOK {
			testSuccess = false
			logs += fmt.Sprintf("Content %s \"%s\" is missing\n", content.content, content.path)
		} else if err != nil {
			testSuccess = false
			logs += err.Error()
		}

	}

	if testSuccess {
		return data.RepositoryReport{
			TestName:    checker.testType,
			GithubRepo:  repoName,
			CheckStatus: data.Successful,
		}
	} else {
		log.Infof("helmchart test failed on repo %s", repoName)
		return data.RepositoryReport{
			TestName:    checker.testType,
			GithubRepo:  repoName,
			CheckStatus: data.Failed,
			Log:         logs,
		}
	}

}
