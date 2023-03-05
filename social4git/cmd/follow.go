package cmd

import (
	"github.com/petar/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	followCmd = &cobra.Command{
		Use:   "follow",
		Short: "Follow a user",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			proto.Follow(ctx, setup.Home, proto.MustParseHandle(ctx, followHandle))
		},
	}
)

var (
	followHandle string
)

func init() {
	rootCmd.AddCommand(followCmd)
	followCmd.Flags().StringVar(&followHandle, "handle", "", "user handle to follow")
	followCmd.MarkFlagRequired("handle")
}
