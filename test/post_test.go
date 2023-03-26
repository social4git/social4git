package test

import (
	"fmt"
	"testing"

	"github.com/gov4git/lib4git/testutil"
	"github.com/social4git/social4git/proto"
)

func TestPost(t *testing.T) {
	ctx := testutil.NewCtx()
	testNet := NewTestNet(ctx, t, 1)
	fmt.Println("testnet", testNet.Dir())

	proto.Post(ctx, testNet.Home(0), []byte("post1"))
	proto.Post(ctx, testNet.Home(0), []byte("post2"))
}
