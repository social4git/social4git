package test

import (
	"fmt"
	"testing"

	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/testutil"
	"github.com/social4git/social4git/proto"
)

// XXX: test with and without cache
// XXX: test local and network repos
func TestSync(t *testing.T) {
	ctx := testutil.NewCtx()
	testNet := NewTestNet(ctx, t, 2)
	fmt.Println("testnet", testNet.Dir())

	proto.Post(ctx, testNet.Home(0), []byte("post1"))
	proto.Post(ctx, testNet.Home(0), []byte("post2"))
	proto.Follow(ctx, testNet.Home(1), testNet.Handle(0))
	proto.Sync(ctx, testNet.Home(1))

	r1 := git.CloneAll(ctx, testNet.Home(1).PrivateReadWrite())
	ft := git.GetBranchTree(ctx, r1.Repo(), proto.PrivateBranch)
	if !FindFileWithContent(ctx, ft, "post1") {
		t.Errorf("expecting to find post")
	}

	// testutil.Hang()
}
