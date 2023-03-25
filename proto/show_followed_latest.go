package proto

import (
	"context"
	"time"
)

func FetchFollowedLatestPostsDay(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchFollowedPostsDay(ctx, home, time.Now().UTC())
}

func FetchFollowedLatestPostsMonth(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchFollowedPostsMonth(ctx, home, time.Now().UTC())
}

func FetchFollowedLatestPostsYear(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchFollowedPostsYear(ctx, home, time.Now().UTC())
}
