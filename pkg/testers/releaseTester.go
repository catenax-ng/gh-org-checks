package testers

import (
	"context"
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
		TestProperty: TestProperty{testName: "Release"},
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
		}}
}

func (tester ReleaseTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	result := tester.ContentTester.PerformTest(repoName, testName)

	release, resp, err := tester.githubClient.Repositories.GetLatestRelease(tester.ctx, tester.owner, repoName)

	if release == nil || resp.StatusCode == http.StatusNotFound {
		result.TestSucceed = result.TestSucceed || false
		result.Log = append(result.Log, "No releases found!\n")
		return result
	} else if err != nil {
		result.TestSucceed = result.TestSucceed || false
		result.Log = append(result.Log, err.Error()+"\n")
		return result
	}

	_, err = semver.StrictNewVersion(*release.Name)

	if err != nil {
		result.TestSucceed = false
		result.Log = append(result.Log, "Not Semantic versioned!\n")
		return result
	}

	result.TestSucceed = result.TestSucceed || true

	return result
}
