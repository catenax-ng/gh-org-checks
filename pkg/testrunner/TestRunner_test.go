package testrunner

import (
	"github.com/google/go-github/v45/github"
	"testing"
)

func Test_listOfOnlyPublicRepositoriesRemainsUnchangedWhenRemovingArchivedOnes(t *testing.T) {
	repositories := []*github.Repository{
		{
			Name:     github.String("just-a-public-example"),
			Archived: github.Bool(false),
		},
		{
			Name:     github.String("just-another-public-example"),
			Archived: github.Bool(false),
		},
	}
	initialRepositoryCount := len(repositories)

	repositories = removeArchivedRepos(repositories)

	if len(repositories) != initialRepositoryCount {
		t.Errorf("expeced length of %v, but got %v", initialRepositoryCount, len(repositories))
	}
}

func Test_shouldRemoveArchivedRepositoriesAndKeepPublicOnes(t *testing.T) {
	repositories := []*github.Repository{
		{
			Name:     github.String("just-a-public-example"),
			Archived: github.Bool(false),
		},
		{
			Name:     github.String("an-archived-repo"),
			Archived: github.Bool(true),
		},
		{
			Name:     github.String("just-another-public-example"),
			Archived: github.Bool(false),
		},
	}

	repositories = removeArchivedRepos(repositories)

	if len(repositories) != 2 {
		t.Errorf("expeced length of %v, but got %v", 2, len(repositories))
	}

	for _, repo := range repositories {
		if *repo.Archived {
			t.Errorf("filtered result contains repository marked as archived. Repo name %s", *repo.Name)
		}
	}
}
