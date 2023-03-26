package cmd

import (
	"github.com/social4git/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Fetch latest posts from users you follow",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			proto.Sync(ctx, setup.Home)
		},
	}
)

func init() {
	rootCmd.AddCommand(syncCmd)
}
