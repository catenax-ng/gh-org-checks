package testers

import (
	"context"
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
		testType:     "ReleaseCheck",
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
	result := tester.contentTest(repoName)

	releases, _, err := tester.githubClient.Repositories.ListReleases(tester.ctx, tester.owner, repoName, &github.ListOptions{})

	if len(releases) == 0 {
		if result.CheckStatus == data.Successful {
			result.CheckStatus = data.Failed
		}

		result.Log += "No releases found!\n"
	} else if err != nil {
		if result.CheckStatus == data.Successful {
			result.CheckStatus = data.Failed
		}

		result.Log += err.Error() + "\n"

	} else {
		result.CheckStatus = data.Successful
	}

	return result
}
