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
	ContentTester
}

func NewReleaseTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new release tester")
	return ReleaseTester{ContentTester{
		testType:     "Release",
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

func (tester ReleaseTester) PerformTest(repoName string) data.RepositoryReport {
	log.Infof("perform release test on repo %s", repoName)
	result := tester.ContentTester.PerformTest(repoName)

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
