package action

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func Set(release string) error {
	localReleases, err := GetLocalReleases()
	if err != nil {
		return nil
	}
	if !slices.Contains(localReleases, release) {
		return fmt.Errorf("version '%s' is not installed", release)
	}
	
	var filePath string
	for _, name := range []string{"ignite", "starport"} {
		p := filepath.Join(getReleasePath(), release, name)
		_, err = os.Stat(p)
		if err == nil {
			filePath = p
			break
		}
		fmt.Println(err)
	}
	if filePath == "" {
		return fmt.Errorf("no binary found in asset")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	binDir := filepath.Join(home, ".local/bin")
	linkPath := filepath.Join(binDir, "ignite")
	fmt.Printf("Make sure %s is in your path!\n", binDir)
	err = os.MkdirAll(binDir, os.ModePerm)
	if err != nil {
		return err
	}
	os.Remove(linkPath)
	err = os.Symlink(
		filePath,
		linkPath,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Set ignite version to %s\n", release)
	
	return nil
}
