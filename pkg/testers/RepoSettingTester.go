package testers

import (
	"context"
	"github.com/catena-x/gh-org-checks/pkg/common"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type RepoSettingTester struct {
	ctx          context.Context
	owner        string
	githubClient *github.Client
}

func NewRepoSettingTester(ctx context.Context, owner string, githubClient *github.Client) RepoSettingTester {
	log.Printf("creating new readme tester")
	return RepoSettingTester{
		ctx:          ctx,
		owner:        owner,
		githubClient: githubClient,
	}
}

func (checker RepoSettingTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	repository, _, _ := checker.githubClient.Repositories.Get(checker.ctx, checker.owner, repoName)

	testResult := data.RepositoryReport{
		TestName:    testName,
		GithubRepo:  repoName,
		TestSucceed: true,
		Log:         []data.LogElement{},
	}

	if repository.GetDefaultBranch() != "main" {
		testResult.TestSucceed = false
		testResult.Log = append(testResult.Log, data.LogElement{
			LogContent: "default Branch is not main",
			RefLink:    common.GetRefLink(common.TrgTwoOne),
		})
	}

	return testResult
}
