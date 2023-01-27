package testers

import (
	"context"

	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type SecurityActionTester struct {
	TestProperty
	ContentTester
}

func NewSecurityActionTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new security github actions tester")
	return SecurityActionTester{
		TestProperty: NewTestProperty("Security Action"),
		ContentTester: NewContentTester(ctx, owner, githubClient,
			[]repositoryContent{
				{
					path:        ".github/workflows",
					contentType: Directory,
				},
				{
					path:        ".github/workflows/veracode.yaml",
					contentType: File,
				},
				{
					path:        ".github/workflows/trivy.yml",
					contentType: File,
				},
				{
					path:        ".github/workflows/kics.yml",
					contentType: File,
				},
			}),
	}
}
