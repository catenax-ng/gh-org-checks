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
	path        string
	contentType contentType
}

type ContentTester struct {
	ctx          context.Context
	owner        string
	githubClient *github.Client
	contents     []repositoryContent
}

func (checker ContentTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	log.Infof("perform content check for %s test on repo %s", testName, repoName)
	var testSuccess = true
	var logs []string

	for _, content := range checker.contents {
		_, _, resp, err := checker.githubClient.Repositories.GetContents(checker.ctx, checker.owner, repoName, content.path, &github.RepositoryContentGetOptions{})

		if resp.StatusCode != http.StatusOK {
			testSuccess = false
			logs = append(logs, fmt.Sprintf("Content %s \"%s\" is missing!", content.contentType, content.path))
		} else if err != nil {
			testSuccess = false
			logs = append(logs, err.Error())
		}

	}

	if testSuccess {
		return data.RepositoryReport{
			TestName:    testName,
			GithubRepo:  repoName,
			TestSucceed: true,
		}
	} else {
		log.Infof("%s test failed on repo %s", testName, repoName)
		return data.RepositoryReport{
			TestName:    testName,
			GithubRepo:  repoName,
			TestSucceed: false,
			Log:         logs,
		}
	}

}
