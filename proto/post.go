package proto

import (
	"context"
	"time"

	"github.com/gov4git/lib4git/git"
)

func Post(
	ctx context.Context,
	home Home,
	content []byte,
) git.Change[PostID] {

	cloned := git.CloneOne(ctx, home.TimelineReadWrite())
	chg := PostLocal(ctx, home, cloned, content)
	cloned.Push(ctx)
	return chg
}

func PostLocal(
	ctx context.Context,
	home Home,
	clone git.Cloned,
	content []byte,
) git.Change[PostID] {

	chg := PostStageOnly(ctx, home, clone, content)
	Commit(ctx, clone.Tree(), chg.Msg)
	return chg
}

func PostStageOnly(
	ctx context.Context,
	home Home,
	clone git.Cloned,
	content []byte,
) git.Change[PostID] {

	postID := NewPostID(time.Now(), content)
	postNS := postID.NS()
	meta := PostMeta{By: home.Handle, ID: postID}
	git.StringToFileStage(ctx, clone.Tree(), postNS.Ext(RawExt), string(content))
	git.ToFileStage(ctx, clone.Tree(), postNS.Ext(MetaExt).Path(), meta)
	return git.Change[PostID]{
		Result: postID,
		Msg:    "post",
	}
}
