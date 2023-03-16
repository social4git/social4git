package proto

import (
	"context"
	"strings"
	"time"

	"github.com/gov4git/lib4git/base"
	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/must"
)

func GetTimelinePostByID(
	ctx context.Context,
	home Home,
	postID PostID,
) PostWithMeta {

	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return GetTimelinePostByIDLocal(ctx, cloned, postID)
}

func GetTimelinePostByIDLocal(
	ctx context.Context,
	clone git.Cloned,
	postID PostID,
) PostWithMeta {

	postNS := postID.NS()
	meta := git.FromFile[PostMeta](ctx, clone.Tree(), postNS.Ext(MetaExt).Path())
	content := git.FileToString(ctx, clone.Tree(), postNS.Ext(RawExt))
	return PostWithMeta{Content: []byte(content), Meta: meta}
}

func ListTimelinePostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return ListTimelinePostsDayLocal(ctx, cloned, day)
}

func ListTimelinePostsDayLocal(
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

func ListTimelinePostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return ListTimelinePostsMonthLocal(ctx, cloned, day)
}

func ListTimelinePostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, m, _ := day.Date()
	for t := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC); t.Month() == m; t = t.AddDate(0, 0, 1) {
		for k := range ListTimelinePostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func ListTimelinePostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) map[PostID]bool {

	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return ListTimelinePostsYearLocal(ctx, cloned, day)
}

func ListTimelinePostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) map[PostID]bool {

	posts := map[PostID]bool{}
	y, _, _ := day.Date()
	for t := time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC); t.Year() == y; t = t.AddDate(0, 0, 1) {
		for k := range ListTimelinePostsDayLocal(ctx, clone, t) {
			posts[k] = true
		}
	}
	return posts
}

func FetchTimelinePostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return FetchTimelinePostsDayLocal(ctx, cloned, day)
}

func FetchTimelinePostsDayLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchTimelinePosts(ctx, clone, ListTimelinePostsDayLocal(ctx, clone, day))
}

func FetchTimelinePostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return FetchTimelinePostsMonthLocal(ctx, cloned, day)
}

func FetchTimelinePostsMonthLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchTimelinePosts(ctx, clone, ListTimelinePostsMonthLocal(ctx, clone, day))
}

func FetchTimelinePostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {

	cloned := git.CloneOne(ctx, home.TimelineReadOnly())
	return FetchTimelinePostsYearLocal(ctx, cloned, day)
}

func FetchTimelinePostsYearLocal(
	ctx context.Context,
	clone git.Cloned,
	day time.Time,
) []PostWithMeta {

	return fetchTimelinePosts(ctx, clone, ListTimelinePostsYearLocal(ctx, clone, day))
}

func fetchTimelinePosts(
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
		posts = append(posts, GetTimelinePostByIDLocal(ctx, clone, postID))
	}
	return posts
}
