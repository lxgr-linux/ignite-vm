package action

import (
	"context"
	"os"

	"github.com/google/go-github/v56/github"
)

type RemoteRelease struct {
	Name string
	Id   int64
	User string
	Repo string
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

	for _, remoteInfo := range [][2]string{{"ignite", "cli"}, {"lxgr-linux", "ignite-cli"}} {
		rels, err := getSpecificRemoteRelease(ctx, ghClient, remoteInfo[0], remoteInfo[1])
		if err != nil {
			return []RemoteRelease{}, err
		}
		releases = append(releases, rels...)
	}
	return
}

func getSpecificRemoteRelease(
	ctx context.Context, ghClient *github.Client, user string, repo string,
) (releases []RemoteRelease, err error) {
	rels, _, err := ghClient.Repositories.ListReleases(ctx, user, repo, nil)
	if err != nil {
		return
	}

	for _, rel := range rels {
		releases = append(releases, RemoteRelease{
			*rel.Name, *rel.ID, user, repo,
		})
	}

	return
}
