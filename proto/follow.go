package proto

import (
	"context"

	"github.com/gov4git/lib4git/git"
)

func Follow(
	ctx context.Context,
	home Home,
	handle Handle,
) git.Change[bool] {

	cloned := git.CloneOne(ctx, home.TimelineReadWrite())
	chg := FollowLocal(ctx, home, cloned, handle)
	cloned.Push(ctx)
	return chg
}

func FollowLocal(
	ctx context.Context,
	home Home,
	clone git.Cloned,
	handle Handle,
) git.Change[bool] {

	chg := FollowStageOnly(ctx, home, clone, handle)
	Commit(ctx, clone.Tree(), chg.Msg)
	return chg
}

func FollowStageOnly(
	ctx context.Context,
	home Home,
	clone git.Cloned,
	handle Handle,
) git.Change[bool] {

	following := GetFollowingLocal(ctx, clone)
	already := following[handle]
	following[handle] = true
	git.ToFileStage(ctx, git.Worktree(ctx, clone.Repo()), FollowingNS.Path(), following)
	return git.Change[bool]{
		Result: !already,
		Msg:    "follow",
	}
}

func GetFollowing(ctx context.Context, home Home) Following {
	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return GetFollowingLocal(ctx, cloned)
}

func GetFollowingLocal(ctx context.Context, clone git.Cloned) Following {
	return git.FromFile[Following](ctx, clone.Tree(), FollowingNS.Path())
}
