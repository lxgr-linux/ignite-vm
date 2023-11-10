package cmd

import (
	"context"
	"fmt"
	"github.com/lxgr-linux/ignite-vm/action"
	"github.com/spf13/cobra"
	"slices"
	"strings"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all awailable ignite versions",
	Long:  "List all awailable ignite versions",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		localReleases, err := action.GetLocalReleases()
		if err != nil {
			err = action.CreateReleaseDir()
			if err != nil {
				return err
			}
		}
		
		remoteReleases, err := action.GetRemoteReleases(ctx)
		if err != nil {
			return err
		}
		
		var versionStrings []string
		for _, rel := range remoteReleases {
			var relString string
			if slices.Contains(localReleases, rel.Name) {
				relString = fmt.Sprintf("%s (installed)", rel.Name)
			} else {
				relString = rel.Name
			}
			versionStrings = append(versionStrings, relString)
		}
		
		fmt.Println(strings.Join(versionStrings, "\n"))
		
		return nil
	},
}

