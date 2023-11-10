package cmd

import (
	"context"
	"github.com/lxgr-linux/ignite-vm/action"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "dowloads and install a certain version of ignite",
	Args: cobra.ExactArgs(1),
	Long:  "dowloads and install a certain version of ignite",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		return action.Install(ctx, args[0])
	},
}