package proto

import (
	"context"
	"path/filepath"
	"time"

	"github.com/gov4git/lib4git/base"
	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/must"
)

func GetFollowingPostByID(
	ctx context.Context,
	home Home,
	postID PostID,
) PostWithMeta {

	cloned := git.CloneOne(ctx, home.FollowingReadOnly())
	return GetFollowingPostByIDLocal(ctx, cloned, postID)
}

func GetFollowingPostByIDLocal(
	ctx context.Context,
	clone git.Cloned,
	postID PostID,
) PostWithMeta {

	postNS := postID.NS()
	meta := git.FromFile[PostMeta](ctx, clone.Tree(), postNS.Ext(MetaExt).Path())
	content := git.FileToString(ctx, clone.Tree(), postNS.Ext(RawExt))
	return PostWithMeta{Content: []byte(content), Meta: meta}
}

func ListFollowingPostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.FollowingReadOnly())
	return ListFollowingPostsDayLocal(ctx, cloned, day)
}

func ListFollowingPostsDayLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	fs := clone.Tree().Filesystem
	dayNS := PostDayNS(day)
	infos, err := fs.ReadDir(dayNS.Path())
	must.NoError(ctx, err)

	posts := map[PostID]bool{}
	for _, info := range infos {
		filename := filepath.Ext(info.Name())
		unparsedID := filename[:len(filename)-len(filepath.Ext(filename))]
		postID, err := ParsePostID(unparsedID)
		if err != nil {
			base.Infof("unrecognized file %s in post directory", filename)
			continue
		}
		posts[postID] = true
	}

	return posts
}

func ListFollowingPostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.FollowingReadOnly())
	return ListFollowingPostsMonthLocal(ctx, cloned, day)
}

func ListFollowingPostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, m, _ := day.Date()
	for t := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC); t.Month() == m; t = t.AddDate(0, 0, 1) {
		for k := range ListFollowingPostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func ListFollowingPostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.FollowingReadOnly())
	return ListFollowingPostsYearLocal(ctx, cloned, day)
}

func ListFollowingPostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, _, _ := day.Date()
	for t := time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC); t.Year() == y; t = t.AddDate(0, 0, 1) {
		for k := range ListFollowingPostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func FetchFollowingPostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.FollowingReadOnly())
	return FetchFollowingPostsDayLocal(ctx, cloned, day)
}

func FetchFollowingPostsDayLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchFollowingPosts(ctx, clone, ListFollowingPostsDayLocal(ctx, clone, day))
}

func FetchFollowingPostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.FollowingReadOnly())
	return FetchFollowingPostsMonthLocal(ctx, cloned, day)
}

func FetchFollowingPostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchFollowingPosts(ctx, clone, ListFollowingPostsMonthLocal(ctx, clone, day))
}

func FetchFollowingPostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.FollowingReadOnly())
	return FetchFollowingPostsYearLocal(ctx, cloned, day)
}

func FetchFollowingPostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchFollowingPosts(ctx, clone, ListFollowingPostsYearLocal(ctx, clone, day))
}

func fetchFollowingPosts(
	ctx context.Context,
	clone git.Cloned,
	postIDMap map[PostID]bool,
) []PostWithMeta {

	postIDs := PostIDs{}
	for postID := range postIDMap {
		postIDs = append(postIDs, postID)
	}
	postIDs.Sort()
	posts := []PostWithMeta{}
	for _, postID := range postIDs {
		posts = append(posts, GetFollowingPostByIDLocal(ctx, clone, postID))
	}
	return posts
}
