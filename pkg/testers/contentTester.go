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
	testType     string
	ctx          context.Context
	owner        string
	githubClient *github.Client
	contents     []repositoryContent
}

func (checker ContentTester) contentTest(repoName string) data.RepositoryReport {

	var testSuccess = true
	var logs = ""

	for _, content := range checker.contents {
		_, _, resp, err := checker.githubClient.Repositories.GetContents(checker.ctx, checker.owner, repoName, content.path, &github.RepositoryContentGetOptions{})

		if resp.StatusCode != http.StatusOK {
			testSuccess = false
			logs += fmt.Sprintf("Content %s \"%s\" is missing!\n", content.contentType, content.path)
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
		log.Infof("%s test failed on repo %s", checker.testType, repoName)
		return data.RepositoryReport{
			TestName:    checker.testType,
			GithubRepo:  repoName,
			CheckStatus: data.Failed,
			Log:         logs,
		}
	}

}
