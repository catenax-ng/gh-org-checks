package testers

import (
	"context"

	"github.com/catena-x/gh-org-checks/pkg/common"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/google/go-github/v45/github"
	log "github.com/sirupsen/logrus"
)

type TRGTwoTester struct {
	TestProperty
	ContentTester
	RepoSettingTester
}

func NewTRGTwoTester(ctx context.Context, owner string, githubClient *github.Client) GithubTester {
	log.Printf("creating new TRG 1 tester")
	return TRGTwoTester{
		TestProperty: NewTestProperty(common.TestTrgTwo),
		ContentTester: NewContentTester(ctx, owner, githubClient,
			[]repositoryContent{
				{
					path:        "docs",
					contentType: Directory,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "charts",
					contentType: Directory,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "AUTHORS.md",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "CODE_OF_CONDUCT.md",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "CONTRIBUTING.md",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "DEPENDENCIES",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "LICENSE",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "NOTICE.md",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "README.md",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "INSTALL.md",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
				{
					path:        "SECURITY.md",
					contentType: File,
					refLink:     common.GetRefLink(common.TrgTwoThree),
				},
			}),
		RepoSettingTester: NewRepoSettingTester(ctx, owner, githubClient),
	}
}

func (tester TRGTwoTester) PerformTest(repoName string, testName string) data.RepositoryReport {
	contentCheckResult := tester.ContentTester.PerformTest(repoName, testName)
	repoCheckResult := tester.RepoSettingTester.PerformTest(repoName, testName)

	return common.MergeReports(contentCheckResult, repoCheckResult)

}
