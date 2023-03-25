package proto

import (
	"context"
)

func FetchLink(
	ctx context.Context,
	link Link,
) PostWithMeta {

	return GetPublishedPostByID(ctx, link.Home(), link.PostID)
}
