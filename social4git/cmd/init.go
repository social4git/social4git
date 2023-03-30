package cmd

import (
	"fmt"
	"os"

	"github.com/gov4git/lib4git/form"
	"github.com/gov4git/lib4git/id"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize your identity",
		Long:  `Generate public and private signing keys and place them in your public and private repositores, respectively.`,
		Run: func(cmd *cobra.Command, args []string) {
			chg := id.Init(ctx, setup.Home.OwnerAddress())
			fmt.Fprint(os.Stdout, form.Pretty(chg.Result))
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}
