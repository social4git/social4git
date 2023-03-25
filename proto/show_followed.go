package proto

import (
	"context"
	"strings"
	"time"

	"github.com/gov4git/lib4git/base"
	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/must"
)

func GetFollowedPostByID(
	ctx context.Context,
	home Home,
	postID PostID,
) PostWithMeta {

	cloned := git.CloneOne(ctx, home.PrivateReadOnly())
	return GetFollowedPostByIDLocal(ctx, cloned, postID)
}

func GetFollowedPostByIDLocal(
	ctx context.Context,
	clone git.Cloned,
	postID PostID,
) PostWithMeta {

	postNS := postID.NS()
	meta := git.FromFile[PostMeta](ctx, clone.Tree(), postNS.Ext(MetaExt).Path())
	content := git.FileToString(ctx, clone.Tree(), postNS.Ext(RawExt))
	return PostWithMeta{Content: []byte(content), Meta: meta}
}

func ListFollowedPostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.PrivateReadOnly())
	return ListFollowedPostsDayLocal(ctx, cloned, day)
}

func ListFollowedPostsDayLocal(
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
		filename := info.Name()
		if !strings.HasSuffix(filename, "."+RawExt) {
			continue
		}
		unparsedID := filename[:len(filename)-len("."+RawExt)]
		postID, err := ParsePostID(unparsedID)
		if err != nil {
			base.Infof("unrecognized file %s in post directory", filename)
			continue
		}
		posts[postID] = true
	}

	return posts
}

func ListFollowedPostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.PrivateReadOnly())
	return ListFollowedPostsMonthLocal(ctx, cloned, day)
}

func ListFollowedPostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, m, _ := day.Date()
	for t := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC); t.Month() == m; t = t.AddDate(0, 0, 1) {
		for k := range ListFollowedPostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func ListFollowedPostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.PrivateReadOnly())
	return ListFollowedPostsYearLocal(ctx, cloned, day)
}

func ListFollowedPostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, _, _ := day.Date()
	for t := time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC); t.Year() == y; t = t.AddDate(0, 0, 1) {
		for k := range ListFollowedPostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func FetchFollowedPostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.PrivateReadOnly())
	return FetchFollowedPostsDayLocal(ctx, cloned, day)
}

func FetchFollowedPostsDayLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchFollowedPosts(ctx, clone, ListFollowedPostsDayLocal(ctx, clone, day))
}

func FetchFollowedPostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.PrivateReadOnly())
	return FetchFollowedPostsMonthLocal(ctx, cloned, day)
}

func FetchFollowedPostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchFollowedPosts(ctx, clone, ListFollowedPostsMonthLocal(ctx, clone, day))
}

func FetchFollowedPostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.PrivateReadOnly())
	return FetchFollowedPostsYearLocal(ctx, cloned, day)
}

func FetchFollowedPostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchFollowedPosts(ctx, clone, ListFollowedPostsYearLocal(ctx, clone, day))
}

func fetchFollowedPosts(
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
		posts = append(posts, GetFollowedPostByIDLocal(ctx, clone, postID))
	}
	return posts
}
