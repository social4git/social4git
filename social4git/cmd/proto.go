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

const (
	DateLayout = "01/02/2006"
)

type Setup struct {
	Home proto.Home
}

type Config struct {
	Handle string `json:"handle"` // e.g. https://github.com/petar/myskrit
	//
	PublicURL  git.URL `json:"public_url"`
	PrivateURL git.URL `json:"private_url"`
	//
	PublicAuth  AuthConfig `json:"public_auth"`
	PrivateAuth AuthConfig `json:"private_auth"`
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

	setAuth(ctx, cfg.PublicAuth, cfg.PublicURL)
	setAuth(ctx, cfg.PrivateAuth, cfg.PrivateURL)

	handle, err := proto.ParseHandle(string(cfg.Handle))
	must.NoError(ctx, err)
	return Setup{
		Home: proto.Home{
			Handle:     handle,
			PublicURL:  cfg.PublicURL,
			PrivateURL: cfg.PrivateURL,
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
