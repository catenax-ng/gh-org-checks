package testers

import (
	"github.com/catena-x/gh-org-checks/pkg/data"
)

type GithubTester interface {
	GetTestName() string
	PerformTest(repoName string, testName string) data.RepositoryReport
}
