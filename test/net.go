package test

import (
	"context"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/gov4git/lib4git/testutil"
	"github.com/petar/social4git/proto"
)

type TestNet struct {
	dir  string
	user []*TestUser
}

func NewTestNet(ctx context.Context, t *testing.T, n int) *TestNet {
	return NewTestNetDir(ctx, t, filepath.Join(t.TempDir(), testutil.UniqueString(ctx)), n)
}

func NewTestNetDir(ctx context.Context, t *testing.T, dir string, n int) *TestNet {
	users := make([]*TestUser, n)
	for i := 0; i < n; i++ {
		users[i] = NewTestUser(ctx, t, filepath.Join(dir, strconv.Itoa(i)))
	}
	return &TestNet{dir: dir, user: users}
}

func (x *TestNet) Dir() string {
	return x.dir
}

func (x *TestNet) Home(i int) proto.Home {
	return x.user[i].Home()
}

func (x *TestNet) Handle(i int) proto.Handle {
	return x.user[i].Handle()
}

type TestUser struct {
	dir     string
	public  testutil.LocalAddress
	private testutil.LocalAddress
}

func NewTestUser(ctx context.Context, t *testing.T, dir string) *TestUser {
	return &TestUser{
		dir:     dir,
		public:  testutil.NewLocalAddressDir(ctx, t, filepath.Join(dir, proto.PublicBranch), proto.PublicBranch, true),
		private: testutil.NewLocalAddressDir(ctx, t, filepath.Join(dir, proto.PrivateBranch), proto.PrivateBranch, true),
	}
}

func (x *TestUser) Dir() string {
	return x.dir
}

func (x *TestUser) Handle() proto.Handle {
	return proto.NewHandle("file", "", x.public.Dir())
}

func (x *TestUser) Home() proto.Home {
	return proto.Home{
		Handle:     x.Handle(),
		PublicURL:  x.public.Address().Repo,
		PrivateURL: x.private.Address().Repo,
	}
}
