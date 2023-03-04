package proto

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/gov4git/lib4git/must"
)

type Link struct {
	Handle
	PostID
}

func NewLink(h Handle, id PostID) Link {
	return Link{Handle: h, PostID: id}
}

func (l Link) URL() *url.URL {
	return &url.URL{
		Scheme: ProtocolName + "_" + l.Handle.Scheme,
		Host:   l.Handle.Host,
		Path:   l.Handle.Path + "?post=" + l.PostID.String(),
	}
}

func (l Link) String() string {
	return l.URL().String()
}

func MustParseLink(ctx context.Context, s string) Link {
	l, err := ParseLink(s)
	must.NoError(ctx, err)
	return l
}

func ParseLink(s string) (Link, error) {
	// parse link as url
	u, err := url.Parse(s)
	if err != nil {
		return Link{}, err
	}
	// parse scheme
	if !strings.HasPrefix(u.Scheme, ProtocolName+"_") {
		return Link{}, fmt.Errorf("link scheme not recognized")
	}
	// parse handle
	h := u.Scheme[len(ProtocolName+"_"):] + "://" + u.Host + "/" + strings.TrimLeft(u.Path, "/")
	handle, err := ParseHandle(h)
	if err != nil {
		return Link{}, err
	}
	// parse id
	p := u.Query().Get("post")
	id, err := ParsePostID(p)
	if err != nil {
		return Link{}, err
	}
	return Link{
		Handle: handle,
		PostID: id,
	}, nil
}
