package testers

import (
	"context"
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

func (checker ReadmeTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	log.Infof("perform reachme check for test %s on repo %s", testName, repoName)
	_, resp, err := checker.githubClient.Repositories.GetReadme(checker.ctx, checker.owner, repoName, &github.RepositoryContentGetOptions{})

	if resp.StatusCode != http.StatusOK {
		log.Infof("readme test failed on repo %s", repoName)
		log.Infof("status code: %d", resp.StatusCode)
		return data.RepositoryReport{
			TestName:    testName,
			GithubRepo:  repoName,
			TestSucceed: false,
			Log:         []string{"Readme file is missing!"},
		}
	}

	if err != nil {
		log.Infof("readme test failed on repo %s", repoName)
		log.Debugf("error message: %s", err.Error())
		return data.RepositoryReport{
			TestName:    testName,
			GithubRepo:  repoName,
			TestSucceed: false,
			Log:         []string{err.Error()},
		}
	}

	log.Infof("readme test is successful on repo %s", repoName)
	return data.RepositoryReport{
		TestName:    testName,
		GithubRepo:  repoName,
		TestSucceed: true,
	}

}
