package cmd

import (
	"github.com/petar/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	unfollowCmd = &cobra.Command{
		Use:   "unfollow",
		Short: "Unfollow a user",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			proto.Unfollow(ctx, setup.Home, proto.MustParseHandle(ctx, unfollowHandle))
		},
	}
)

var (
	unfollowHandle string
)

func init() {
	rootCmd.AddCommand(unfollowCmd)
	unfollowCmd.Flags().StringVar(&unfollowHandle, "handle", "", "user handle to unfollow")
	unfollowCmd.MarkFlagRequired("handle")
}
