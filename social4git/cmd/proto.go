package cmd

import (
	"context"

	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/must"
	"github.com/petar/social4git/proto"
)

const (
	AgentName           = "social4git"
	AgentVarPath        = "." + AgentName
	AgentConfigFilebase = "config.json"
	AgentTempPath       = AgentName
)

type Setup struct {
	Home proto.Home
}

type Config struct {
	Handle string `json:"handle"` // e.g. https://github.com/petar/myskrit
	//
	TimelineURL  git.URL `json:"timeline_url"`
	FollowingURL git.URL `json:"following_url"`
	//
	TimelineAuth  AuthConfig `json:"timeline_auth"`
	FollowingAuth AuthConfig `json:"following_auth"`
	//
	VarDir string `json:"var_dir"`
}

type AuthConfig struct {
	SSHPrivateKeysFile *string       `json:"ssh_private_keys_file"`
	AccessToken        *string       `json:"access_token"`
	UserPassword       *UserPassword `json:"user_password"`
}

type UserPassword struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (cfg Config) Setup(ctx context.Context) Setup {

	git.SetAuthor(proto.ProtocolName+" agent", "no-reply@"+proto.ProtocolName+".xyz")

	setAuth(ctx, cfg.TimelineAuth, cfg.TimelineURL)
	setAuth(ctx, cfg.FollowingAuth, cfg.FollowingURL)

	handle, err := proto.ParseHandle(string(cfg.Handle))
	must.NoError(ctx, err)
	return Setup{
		Home: proto.Home{
			Handle:       handle,
			TimelineURL:  cfg.TimelineURL,
			FollowingURL: cfg.FollowingURL,
		},
	}
}

func setAuth(ctx context.Context, authConfig AuthConfig, url git.URL) {
	switch {
	case authConfig.SSHPrivateKeysFile != nil:
		git.SetAuth(ctx, url, git.MakeSSHFileAuth(ctx, "git", *authConfig.SSHPrivateKeysFile))
	case authConfig.AccessToken != nil:
		git.SetAuth(ctx, url, git.MakeTokenAuth(ctx, *authConfig.AccessToken))
	case authConfig.UserPassword != nil:
		git.SetAuth(ctx, url, git.MakePasswordAuth(ctx, authConfig.UserPassword.User, authConfig.UserPassword.Password))
	}
}
