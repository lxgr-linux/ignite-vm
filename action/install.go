package action

import (
	"context"
	"fmt"
	"github.com/google/go-github/v56/github"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func Install(ctx context.Context, release string) error {
	fmt.Printf("Installing ignite %s...\n", release)
	arch := getArch()
	ops, err := getOS()
	if err != nil {
		return err
	}
	
	client := getClient()
	remoteReleases, err := GetRemoteReleases(ctx)
	if err != nil {
		return err
	}
	
	var releaseObj *RemoteRelease
	for _, rel := range remoteReleases {
		if rel.Name == release {
			releaseObj = &rel
			break
		}
	}
	
	if releaseObj == nil {
		return fmt.Errorf(
			"release '%s' ist not awailable; see 'ignite-vm list' to get all awailable releases",
			release,
		)
	}
	
	assets, _, err := client.Repositories.ListReleaseAssets(ctx, "ignite", "cli", releaseObj.Id, nil)
	if err != nil {
		return err
	}
	
	for _, asset := range assets {
		info := getInfoFromAsset(asset)
		if info.Arch == arch && info.Os == ops {
			releaseDir := filepath.Join(getReleasePath(), "/"+release)
			dlPath := filepath.Join(releaseDir, "/dl.tar.gz")
			err = os.MkdirAll(releaseDir, os.ModePerm)
			if err != nil {
				return err
			}
			err = downloadAsset(dlPath, *asset.BrowserDownloadURL)
			if err != nil {
				return err
			}
			err = unpackDl(dlPath, releaseDir)
			if err != nil {
				return err
			}
			
			return nil
		}
	}
	
	return fmt.Errorf("error while finding proper asset")
}

type ReleaseInfo struct {
	Raw *github.ReleaseAsset
	Os string
	Arch string
}

func getInfoFromAsset(asset *github.ReleaseAsset) ReleaseInfo {
	splid := strings.Split(strings.Trim(*asset.Name, ".tar.gz"), "_")
	
	return ReleaseInfo{
		Raw: asset,
		Arch: splid[len(splid) - 1],
		Os: splid[len(splid) - 2],
	}
}

func getOS() (os string, err error) {
	ops := runtime.GOOS
	if slices.Contains([]string{"linux", "darwin"}, ops) {
		os = ops
	} else {
		err = fmt.Errorf("platform '%s' is not supported", ops)
	}
	return 
}

func getArch() string {
	return runtime.GOARCH
}

func downloadAsset(path string, url string) (err error) {
	fmt.Println("Downloading asset...")

	// Create the file
	out, err := os.Create(path)
	if err != nil  {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		return err
	}

	return nil
}

func unpackDl(archivePath string, outputDir string) error {
	fmt.Println("Extracting asset...")
	cmd := exec.Command(
		"tar",
		"-xf",
		archivePath,
		"-C",
		outputDir,
		)
	return cmd.Run()
}
