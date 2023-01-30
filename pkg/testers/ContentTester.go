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
	refLink     string
	contentType contentType
}

type ContentTester struct {
	ctx          context.Context
	owner        string
	githubClient *github.Client
	contents     []repositoryContent
}

func NewContentTester(ctx context.Context, owner string, githubClient *github.Client, contents []repositoryContent) ContentTester {
	log.Printf("creating new content tester")
	return ContentTester{
		ctx:          ctx,
		owner:        owner,
		githubClient: githubClient,
		contents:     contents,
	}
}

func (checker ContentTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	log.Infof("perform content check for %s test on repo %s", testName, repoName)
	var testSuccess = true
	var logItem []data.LogElement

	for _, content := range checker.contents {
		_, _, resp, err := checker.githubClient.Repositories.GetContents(checker.ctx, checker.owner, repoName, content.path, &github.RepositoryContentGetOptions{})

		if resp.StatusCode != http.StatusOK {
			testSuccess = false
			logItem = append(logItem, data.LogElement{
				LogContent: fmt.Sprintf("Content %s \"%s\" is missing!", content.contentType, content.path),
				RefLink:    content.refLink,
			})
		} else if err != nil {
			testSuccess = false
			logItem = append(logItem, data.LogElement{
				LogContent: err.Error(),
				RefLink:    content.refLink,
			})
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
			Log:         logItem,
		}
	}

}
