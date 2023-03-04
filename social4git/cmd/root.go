package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gov4git/lib4git/base"
	"github.com/gov4git/lib4git/form"
	"github.com/gov4git/lib4git/git"
	"github.com/petar/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   proto.ProtocolName,
		Short: proto.ProtocolName + " is a command-line client for a decentralized Twitter-like app over git",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

var ctx = git.WithAuth(context.Background(), nil)

var (
	configPath string
	verbose    bool
)

func init() {
	cobra.OnInitialize(initAfterFlags)

	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", defaultConfigPath, "config file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "run in developer mode with verbose logging")
}

var defaultConfigPath = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		base.Fatalf("looking for home dir (%v)", err)
	}
	base.AssertNoErr(err)
	return filepath.Join(home, AgentVarPath, AgentConfigFilebase)
}()

func initAfterFlags() {
	if verbose {
		base.LogVerbosely()
	} else {
		base.LogQuietly()
	}

	if configPath == "" {
		configPath = defaultConfigPath
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		base.Fatalf("reading config file (%v)", err)
	}

	config, err := form.DecodeBytes[Config](ctx, data)
	if err != nil {
		base.Fatalf("decoding config file (%v)", err)
	}

	if config.VarDir != "" {
		git.UseCache(ctx, filepath.Join(config.VarDir, "cache"))
	}

	setup = config.Setup(ctx)
}

var (
	setup Setup
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
