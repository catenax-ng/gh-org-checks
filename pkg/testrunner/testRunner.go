package testrunner

import (
	"context"
	"github.com/catena-x/gh-org-checks/pkg/data"
	"github.com/catena-x/gh-org-checks/pkg/testers"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
	"log"
	"os"
	"time"
)

type fn func(ctx context.Context, owner string, githubClient *github.Client) testers.GithubTester

type TestRunner struct {
	ctx        context.Context
	githubOrg  string
	client     *github.Client
	testSuites []testers.GithubTester
}

func NewTestRunner() *TestRunner {
	githubOrg := os.Getenv("GITHUB_ORG_NAME")
	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")

	if githubOrg == "" || githubAccessToken == "" {
		panic("Missing essential environment variable: GITHUB_ORG_NAME, GITHUB_ACCESS_TOKEN")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return &TestRunner{
		ctx:        ctx,
		githubOrg:  githubOrg,
		client:     client,
		testSuites: []testers.GithubTester{},
	}
}

func (runner *TestRunner) AddToTestSuites(f fn) {
	githubTester := f(runner.ctx, runner.githubOrg, runner.client)
	runner.testSuites = append(runner.testSuites, githubTester)

}

func (runner *TestRunner) PerformRepoChecks() data.OrgReports {
	repos, _, err := runner.client.Repositories.ListByOrg(runner.ctx, runner.githubOrg, &github.RepositoryListByOrgOptions{})

	if err != nil {
		return data.OrgReports{
			Error: err,
		}
	}

	orgReport := data.OrgReports{
		OrgName:             runner.githubOrg,
		LastTestTime:        time.Now(),
		RepositoriesReports: []data.RepositoriesReports{},
	}

	for _, repo := range repos {
		repoName := *repo.Name
		log.Printf("Checking repositroy: " + repoName)

		reposReport := data.RepositoriesReports{
			RepositoryName:   repoName,
			RepositoryReport: []data.RepositoryReport{},
		}

		for _, test := range runner.testSuites {
			report := test.PerformTest(repoName)
			reposReport.RepositoryReport = append(reposReport.RepositoryReport, report)
		}

		orgReport.RepositoriesReports = append(orgReport.RepositoriesReports, reposReport)
	}

	return orgReport
}