package proto

import (
	"context"

	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/ns"
)

func Sync(
	ctx context.Context,
	home Home,
) git.ChangeNoResult {

	cloned := git.CloneAll(ctx, home.FollowingReadWrite())
	chg := SyncLocal(ctx, home, cloned)
	cloned.Push(ctx)
	return chg
}

func SyncLocal(
	ctx context.Context,
	home Home,
	clone git.Cloned,
) git.ChangeNoResult {

	following := GetFollowingLocal(ctx, clone)
	addrs := []git.Address{}
	caches := []git.Branch{}
	timelineNS := []ns.NS{}
	for handle := range following {
		u := handle.URL()
		addrs = append(addrs, git.NewAddress(u, TimelineBranch))
		caches = append(caches, git.Branch(CacheBranch(u)))
		timelineNS = append(timelineNS, TimelineNS)
	}

	git.EmbedOnBranch(
		ctx,
		clone.Repo(),
		addrs,
		caches,
		FollowingBranch,
		timelineNS,
		false,
		FilterPosts,
	)

	return git.ChangeNoResult{Msg: "sync"}
}
