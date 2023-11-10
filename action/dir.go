package action

import (
	"github.com/adrg/xdg"
	"os"
	"path/filepath"
)

func getReleasePath() string {
	return filepath.Join(xdg.CacheHome, "/ignite-vm/releases")
}

func CreateReleaseDir() error {
	return os.MkdirAll(getReleasePath(), os.ModePerm)
}