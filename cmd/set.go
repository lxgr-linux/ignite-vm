package cmd

import (
	"github.com/lxgr-linux/ignite-vm/action"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set [version]",
	Short: "Sets an ingite version",
	Args: cobra.ExactArgs(1),
	Long:  "Sets an ingite version",
	RunE: func(cmd *cobra.Command, args []string) error {
		return action.Set(args[0])
	},
}