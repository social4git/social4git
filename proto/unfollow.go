package proto

import (
	"context"

	"github.com/gov4git/lib4git/git"
)

func Unfollow(
	ctx context.Context,
	home Home,
	handle Handle,
) git.Change[bool] {

	cloned := git.CloneOne(ctx, home.PublicReadWrite())
	chg := UnfollowLocal(ctx, home, cloned, handle)
	cloned.Push(ctx)
	return chg
}

func UnfollowLocal(
	ctx context.Context,
	home Home,
	clone git.Cloned,
	handle Handle,
) git.Change[bool] {

	chg := UnfollowStageOnly(ctx, home, clone, handle)
	Commit(ctx, clone.Tree(), chg.Msg)
	return chg
}

func UnfollowStageOnly(
	ctx context.Context,
	home Home,
	clone git.Cloned,
	handle Handle,
) git.Change[bool] {

	following := GetFollowingLocal(ctx, clone)
	already := following[handle]
	delete(following, handle)
	git.ToFileStage(ctx, git.Worktree(ctx, clone.Repo()), FollowingNS.Path(), HandleSetToUnparsedHandleList(following))
	return git.Change[bool]{
		Result: already,
		Msg:    "unfollow",
	}
}
