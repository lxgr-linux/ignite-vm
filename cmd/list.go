package cmd

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/lxgr-linux/ignite-vm/action"
	"github.com/spf13/cobra"
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
			var relString = rel.Name

			if rel.User != "ignite" || rel.Repo != "cli" {
				relString += fmt.Sprintf(" [%s/%s]", rel.User, rel.Repo)
			}

			if slices.Contains(localReleases, rel.Name) {
				relString += " (installed)"
			}

			versionStrings = append(versionStrings, relString)
		}

		fmt.Println(strings.Join(versionStrings, "\n"))

		return nil
	},
}
