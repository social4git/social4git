package proto

import (
	"context"
	"time"
)

func FetchFollowingLatestPostsDay(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchFollowingPostsDay(ctx, home, time.Now().UTC())
}

func FetchFollowingLatestPostsMonth(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchFollowingPostsMonth(ctx, home, time.Now().UTC())
}

func FetchFollowingLatestPostsYear(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchFollowingPostsYear(ctx, home, time.Now().UTC())
}
