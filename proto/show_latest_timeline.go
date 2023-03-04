package proto

import (
	"context"
	"time"
)

func FetchTimelineLatestPostsDay(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {
	return FetchTimelinePostsDay(ctx, home, time.Now().UTC())
}

func FetchTimelineLatestPostsMonth(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {
	return FetchTimelinePostsMonth(ctx, home, time.Now().UTC())
}

func FetchTimelineLatestPostsYear(
	ctx context.Context,
	home Home,
	day time.Time,
) []PostWithMeta {
	return FetchTimelinePostsYear(ctx, home, time.Now().UTC())
}
