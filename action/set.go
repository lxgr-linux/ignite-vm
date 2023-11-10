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
	
	filePath := filepath.Join(getReleasePath(), release, "ignite")
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	binDir := filepath.Join(home, ".local/bin")
	fmt.Printf("Make sure %s is in your path!\n", binDir)
	linkPath := filepath.Join(binDir, "ignite")
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
