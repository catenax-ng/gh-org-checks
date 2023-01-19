package testers

import (
	"context"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type OSSTester struct {
	TestProperty
	ContentTester
}

func NewOSSTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new OSS tester")
	return OSSTester{
		TestProperty: TestProperty{testName: "OSS"},
		ContentTester: ContentTester{
			ctx:          ctx,
			owner:        owner,
			githubClient: githubClient,
			contents: []repositoryContent{
				{
					path:        "LICENSE",
					contentType: File,
				},
				{
					path:        "DEPENDENCIES",
					contentType: File,
				},
				{
					path:        "NOTICE.md",
					contentType: File,
				},
				{
					path:        "SECURITY.md",
					contentType: File,
				},
			},
		}}
}
