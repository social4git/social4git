package cmd

import (
	"fmt"

	"github.com/petar/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	followingCmd = &cobra.Command{
		Use:   "following",
		Short: "List users you are following",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			h := setup.Home.Handle
			if followingHandle != "" {
				h = proto.MustParseHandle(ctx, followingHandle)
			}
			for _, h := range proto.FollowingToHandles(proto.GetFollowing(ctx, h.Home())) {
				fmt.Println(h)
			}
		},
	}
)

var (
	followingHandle string
)

func init() {
	rootCmd.AddCommand(followingCmd)
	followingCmd.Flags().StringVarP(&followingHandle, "handle", "h", "", "user handle")
}
