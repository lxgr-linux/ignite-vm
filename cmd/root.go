package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ignite-vm",
	Short: "Ignite version manager for working seamlessly with different ignite versions",
	Long:  "github.com/lxgr-linux/ignite-vm",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("arguments have to be given")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
