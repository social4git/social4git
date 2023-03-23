package test

import (
	"context"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/gov4git/lib4git/must"
)

func FindFileWithContent(ctx context.Context, t *object.Tree, content string) bool {
	iter := t.Files()
	for {
		f, err := iter.Next()
		if err != nil {
			break
		}
		c, err := f.Contents()
		must.NoError(ctx, err)
		if c == content {
			return true
		}
	}
	return false
}
