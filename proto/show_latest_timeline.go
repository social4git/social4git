package proto

import (
	"context"
	"time"

	"github.com/gov4git/lib4git/base"
)

func FetchTimelineLatestPostsDay(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchTimelinePostsDay(ctx, home, time.Now().UTC())
}

func FetchTimelineLatestPostsMonth(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	base.Infof("FetchTimelineLatestPostsMonth")
	return FetchTimelinePostsMonth(ctx, home, time.Now().UTC())
}

func FetchTimelineLatestPostsYear(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchTimelinePostsYear(ctx, home, time.Now().UTC())
}
