package cmd

import (
	"fmt"
	"os"

	"github.com/social4git/social4git"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(os.Stdout, social4git.Version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
