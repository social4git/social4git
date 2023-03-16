package proto

import (
	"context"
	"os"

	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/must"
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
	git.ToFileStage(ctx, git.Worktree(ctx, clone.Repo()), FollowingNS.Path(), HandleSetToUnparsedHandleList(following))
	return git.Change[bool]{
		Result: !already,
		Msg:    "follow",
	}
}

func GetFollowing(ctx context.Context, home Home) HandleSet {
	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return GetFollowingLocal(ctx, cloned)
}

func GetFollowingLocal(ctx context.Context, clone git.Cloned) HandleSet {
	unparsed, err := git.TryFromFile[UnparsedHandleList](ctx, clone.Tree(), FollowingNS.Path())
	if err == os.ErrNotExist {
		unparsed = UnparsedHandleList{}
	} else {
		must.NoError(ctx, err)
	}
	return UnparsedHandleListToHandleSet(unparsed)
}
