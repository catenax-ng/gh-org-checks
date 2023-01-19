package testers

import (
	"context"
	"github.com/catena-x/gh-org-checks/pkg/common"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type TRGOneTester struct {
	TestProperty
	ContentTester
	ReadmeTester
}

func NewTRGOneTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new TRG 1 tester")
	return TRGOneTester{
		TestProperty: TestProperty{testName: "TRG 1"},
		ContentTester: ContentTester{
			ctx:          ctx,
			owner:        owner,
			githubClient: githubClient,
			contents: []repositoryContent{
				{
					path:        "CHANGELOG.md",
					contentType: File,
				},
			},
		},
		ReadmeTester: ReadmeTester{
			ctx:          ctx,
			owner:        owner,
			githubClient: githubClient,
		},
	}
}

func (tester TRGOneTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	contentCheckResult := tester.ContentTester.PerformTest(repoName, testName)
	readmeCheckResult := tester.ReadmeTester.PerformTest(repoName, testName)

	return common.MergeReports(contentCheckResult, readmeCheckResult)

}
