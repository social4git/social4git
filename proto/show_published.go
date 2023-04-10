package proto

import (
	"context"
	"strings"
	"time"

	"github.com/gov4git/lib4git/base"
	"github.com/gov4git/lib4git/git"
)

func GetPublishedPostByID(
	ctx context.Context,
	home Home,
	postID PostID,
) PostWithMeta {

	cloned := git.CloneOne(ctx, home.PublicReadOnly())
	return GetPublishedPostByIDLocal(ctx, cloned, postID)
}

func GetPublishedPostByIDLocal(
	ctx context.Context,
	clone git.Cloned,
	postID PostID,
) PostWithMeta {

	postNS := postID.NS()
	meta := git.FromFile[PostMeta](ctx, clone.Tree(), postNS.Ext(MetaExt).Path())
	content := git.FileToString(ctx, clone.Tree(), postNS.Ext(RawExt))
	return PostWithMeta{Content: []byte(content), Meta: meta}
}

func ListPublishedPostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.PublicReadOnly())
	return ListPublishedPostsDayLocal(ctx, cloned, day)
}

func ListPublishedPostsDayLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	fs := clone.Tree().Filesystem
	dayNS := PostDayNS(day)
	infos, err := fs.ReadDir(dayNS.Path())
	if err != nil { // no such file or directory
		return nil
	}

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

func ListPublishedPostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.PublicReadOnly())
	return ListPublishedPostsMonthLocal(ctx, cloned, day)
}

func ListPublishedPostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, m, _ := day.Date()
	for t := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC); t.Month() == m; t = t.AddDate(0, 0, 1) {
		for k := range ListPublishedPostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func ListPublishedPostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.PublicReadOnly())
	return ListPublishedPostsYearLocal(ctx, cloned, day)
}

func ListPublishedPostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, _, _ := day.Date()
	for t := time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC); t.Year() == y; t = t.AddDate(0, 0, 1) {
		for k := range ListPublishedPostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func FetchPublishedPostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.PublicReadOnly())
	return FetchPublishedPostsDayLocal(ctx, cloned, day)
}

func FetchPublishedPostsDayLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchPublishedPosts(ctx, clone, ListPublishedPostsDayLocal(ctx, clone, day))
}

func FetchPublishedPostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.PublicReadOnly())
	return FetchPublishedPostsMonthLocal(ctx, cloned, day)
}

func FetchPublishedPostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchPublishedPosts(ctx, clone, ListPublishedPostsMonthLocal(ctx, clone, day))
}

func FetchPublishedPostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.PublicReadOnly())
	return FetchPublishedPostsYearLocal(ctx, cloned, day)
}

func FetchPublishedPostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchPublishedPosts(ctx, clone, ListPublishedPostsYearLocal(ctx, clone, day))
}

func fetchPublishedPosts(
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
		posts = append(posts, GetPublishedPostByIDLocal(ctx, clone, postID))
	}
	return posts
}
