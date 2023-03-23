package test

import (
	"fmt"
	"testing"

	"github.com/gov4git/lib4git/testutil"
	"github.com/petar/social4git/proto"
)

// XXX: test with and without cache
func TestSync(t *testing.T) {
	ctx := testutil.NewCtx()
	testNet := NewTestNet(ctx, t, 2)
	fmt.Println("testnet", testNet.Dir())

	proto.Post(ctx, testNet.Home(0), []byte("post1"))
	proto.Post(ctx, testNet.Home(0), []byte("post2"))

	proto.Follow(ctx, testNet.Home(1), testNet.Handle(0))

	proto.Sync(ctx, testNet.Home(1))

	// testutil.Hang()
}
