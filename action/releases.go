package action

import (
	"context"
	"github.com/google/go-github/v56/github"
	"os"
)

type RemoteRelease struct {
	Name string
	Id int64
}

func getClient() *github.Client {
	return github.NewClient(nil)
}

func GetLocalReleases() (releases []string, err error) {
	dir, err := os.ReadDir(getReleasePath())
	if err != nil {
		return
	}
	for _, sub := range dir {
		if sub.IsDir() {
			releases = append(releases, sub.Name())
		}
	}
	
	return
}

func GetRemoteReleases(ctx context.Context) (releases []RemoteRelease, err error) {
	ghClient := getClient()
	rels, _, err := ghClient.Repositories.ListReleases(ctx, "ignite", "cli", nil)
	if err != nil {
		return
	}
	for _, rel := range rels {
		releases = append(releases, RemoteRelease{*rel.Name, *rel.ID})
	}
	return
}