package proto

import (
	"context"
	"time"
)

func FetchPublishedLatestPostsDay(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchPublishedPostsDay(ctx, home, time.Now().UTC())
}

func FetchPublishedLatestPostsMonth(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchPublishedPostsMonth(ctx, home, time.Now().UTC())
}

func FetchPublishedLatestPostsYear(
	ctx context.Context,
	home Home,
) []PostWithMeta {
	return FetchPublishedPostsYear(ctx, home, time.Now().UTC())
}
