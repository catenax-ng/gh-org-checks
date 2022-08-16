package testers

import (
	"github.com/catena-x/gh-org-checks/pkg/data"
)

type GithubTester interface {
	PerformTest(repoName string) data.RepositoryReport
}
