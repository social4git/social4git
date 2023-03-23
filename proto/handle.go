package proto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/gov4git/lib4git/git"
	"github.com/gov4git/lib4git/must"
)

// example handle : https://example.com:8080/git/repo.git
// example link to post : social4git_https://example.com:8080/git/repo.git?post=20230123231112_abcd_fghi

type Handle struct {
	Scheme string
	Host   string
	Path   string
}

func NewHandle(scheme string, host string, path string) Handle {
	return Handle{Scheme: scheme, Host: host, Path: path}
}

func (h Handle) Home() Home {
	return Home{
		Handle:      h,
		TimelineURL: h.GitURL(),
	}
}

func (h Handle) DebugString() string {
	return fmt.Sprintf("scheme=%v host=%v path=%v", h.Scheme, h.Host, h.Path)
}

func (h Handle) String() string {
	return h.URL().String()
}

func (h Handle) URL() *url.URL {
	return &url.URL{Scheme: h.Scheme, Host: h.Host, Path: "/" + strings.Trim(h.Path, "/")}
}

func (h Handle) GitURL() git.URL {
	return git.URL(h.URL().String())
}

func (h Handle) MarshalJSON() ([]byte, error) {
	s := h.String()
	return json.Marshal(s)
}

func (h *Handle) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	g, err := ParseHandle(s)
	if err != nil {
		return err
	}
	*h = g
	return nil
}

func MustParseHandle(ctx context.Context, urlOrHandle string) Handle {
	h, err := ParseHandle(urlOrHandle)
	must.NoError(ctx, err)
	return h
}

func ParseHandle(s string) (Handle, error) {
	u, err := url.Parse(s)
	if err != nil {
		return Handle{}, err
	}
	if u.Scheme != "https" && u.Scheme != "file" {
		return Handle{}, fmt.Errorf("handle must be an https or file url")
	}
	return Handle{
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   "/" + strings.Trim(u.Path, "/"),
	}, nil
}

type Handles []Handle

func (hs Handles) Len() int {
	return len(hs)
}

func (hs Handles) Less(i, j int) bool {
	return hs[i].String() < hs[j].String()
}

func (hs Handles) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func (hs Handles) Sort() {
	sort.Sort(hs)
}

func FollowingToHandles(f HandleSet) Handles {
	hs := Handles{}
	for h := range f {
		hs = append(hs, h)
	}
	hs.Sort()
	return hs
}
