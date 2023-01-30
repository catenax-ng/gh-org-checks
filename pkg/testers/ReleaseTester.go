package testers

import (
	"context"
	"github.com/catena-x/gh-org-checks/pkg/common"
	"net/http"

	"github.com/Masterminds/semver/v3"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type ReleaseTester struct {
	TestProperty
	ContentTester
}

func NewReleaseTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new release tester")
	return ReleaseTester{
		TestProperty: NewTestProperty("Release"),
		ContentTester: NewContentTester(ctx, owner, githubClient,
			[]repositoryContent{
				{
					path:        "CHANGELOG.md",
					contentType: File,
				},
			}),
	}
}

func (tester ReleaseTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	result := tester.ContentTester.PerformTest(repoName, testName)
	release, resp, err := tester.githubClient.Repositories.GetLatestRelease(tester.ctx, tester.owner, repoName)

	furtherReport := data.RepositoryReport{
		TestName:   testName,
		GithubRepo: repoName,
		Log:        []data.LogElement{},
	}

	if release == nil || resp.StatusCode == http.StatusNotFound {
		furtherReport.TestSucceed = false

		furtherReport.Log = append(furtherReport.Log, data.LogElement{
			LogContent: "No releases found!",
		})
		return common.MergeReports(result, furtherReport)
	} else if err != nil {
		furtherReport.TestSucceed = false
		furtherReport.Log = append(furtherReport.Log, data.LogElement{
			LogContent: err.Error(),
		})

		return common.MergeReports(result, furtherReport)
	}

	_, err = semver.StrictNewVersion(*release.Name)

	if err != nil {
		furtherReport.TestSucceed = false
		furtherReport.Log = append(furtherReport.Log, data.LogElement{
			LogContent: "Not Semantic versioned!",
		})
		return common.MergeReports(result, furtherReport)
	}

	furtherReport.TestSucceed = true
	return common.MergeReports(result, furtherReport)
}
