package testers

import (
	"context"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/google/go-github/v45/github"
	"net/http"
)

type ReadmeTester struct {
	testType     string
	ctx          context.Context
	owner        string
	githubClient *github.Client
}

func NewReadMeTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	return ReadmeTester{
		testType:     "ReadmeCheck",
		ctx:          ctx,
		owner:        owner,
		githubClient: githubClient,
	}
}

func (checker ReadmeTester) PerformTest(repoName string) data.RepositoryReport {
	_, resp, err := checker.githubClient.Repositories.GetReadme(checker.ctx, checker.owner, repoName, &github.RepositoryContentGetOptions{})

	if err != nil {
		return data.RepositoryReport{
			TestName:    checker.testType,
			GithubRepo:  repoName,
			CheckStatus: data.Failed,
			Log:         err.Error(),
		}
	}

	if resp.StatusCode != http.StatusOK {
		return data.RepositoryReport{
			TestName:    checker.testType,
			GithubRepo:  repoName,
			CheckStatus: data.Failed,
			Log:         "Readme file is not present!",
		}
	}

	return data.RepositoryReport{
		TestName:    checker.testType,
		GithubRepo:  repoName,
		CheckStatus: data.Successful,
	}

}
