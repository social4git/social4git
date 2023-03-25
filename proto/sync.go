package proto

import (
	"context"

	"github.com/gov4git/lib4git/base"
	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/ns"
)

func Sync(
	ctx context.Context,
	home Home,
) git.ChangeNoResult {

	cloned := git.CloneAll(ctx, home.PrivateReadWrite())
	chg := SyncLocal(ctx, home, cloned)
	cloned.Push(ctx)
	return chg
}

func SyncLocal(
	ctx context.Context,
	home Home,
	clone git.Cloned,
) git.ChangeNoResult {

	following := GetFollowing(ctx, home)
	addrs := []git.Address{}
	caches := []git.Branch{}
	publicNS := []ns.NS{}
	for handle := range following {
		u := handle.GitURL()
		addrs = append(addrs, git.NewAddress(u, PublicBranch))
		caches = append(caches, git.Branch(CacheBranch(u)))
		publicNS = append(publicNS, PublicNS)
	}

	if len(addrs) == 0 {
		base.Infof("not following anyone")
	} else {
		git.EmbedOnBranch(
			ctx,
			clone.Repo(),
			addrs,
			caches,
			PrivateBranch,
			publicNS,
			false,
			FilterPosts,
		)
	}

	return git.ChangeNoResult{Msg: "sync"}
}
