package proto

import (
	"context"
)

func FetchLink(
	ctx context.Context,
	link Link,
) PostWithMeta {

	return GetTimelinePostByID(ctx, link.Home(), link.PostID)
}
