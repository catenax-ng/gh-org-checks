package testers

import (
	"context"
	"github.com/catena-x/gh-org-checks/pkg/common"
	"net/http"

	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type ReadmeTester struct {
	ctx          context.Context
	owner        string
	githubClient *github.Client
}

func NewReadmeTester(ctx context.Context, owner string, githubClient *github.Client) ReadmeTester {
	log.Printf("creating new readme tester")
	return ReadmeTester{
		ctx:          ctx,
		owner:        owner,
		githubClient: githubClient,
	}
}

func (checker ReadmeTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	log.Infof("perform reachme check for test %s on repo %s", testName, repoName)
	_, resp, err := checker.githubClient.Repositories.GetReadme(checker.ctx, checker.owner, repoName, &github.RepositoryContentGetOptions{})

	logItem := data.LogElement{
		RefLink: common.GetRefLink(common.TrgOneOne),
	}

	if resp.StatusCode != http.StatusOK {
		log.Infof("readme test failed on repo %s", repoName)
		log.Infof("status code: %d", resp.StatusCode)
		logItem.LogContent = "Readme file is missing!"
		return data.RepositoryReport{
			TestName:    testName,
			GithubRepo:  repoName,
			TestSucceed: false,
			Log: []data.LogElement{
				logItem,
			},
		}
	}

	if err != nil {
		log.Infof("readme test failed on repo %s", repoName)
		log.Debugf("error message: %s", err.Error())
		logItem.LogContent = err.Error()
		return data.RepositoryReport{
			TestName:    testName,
			GithubRepo:  repoName,
			TestSucceed: false,
			Log: []data.LogElement{
				logItem,
			},
		}
	}

	log.Infof("readme test is successful on repo %s", repoName)
	return data.RepositoryReport{
		TestName:    testName,
		GithubRepo:  repoName,
		TestSucceed: true,
	}

}
