package cmd

import (
	"fmt"

	"github.com/petar/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "Fetch any post",
		Long:  `Fetch a post from anyone using a link`,
		Run: func(cmd *cobra.Command, args []string) {
			link := proto.MustParseLink(ctx, fetchLink)
			pm := proto.FetchLink(ctx, link)
			fmt.Println(pm)
		},
	}
)

var (
	fetchLink string
)

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVar(&fetchLink, "link", "", "link to post")
}
