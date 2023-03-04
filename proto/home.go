package proto

import "github.com/gov4git/lib4git/git"

type Home struct {
	Handle       Handle
	TimelineURL  git.URL
	FollowingURL git.URL
}

func (h Home) Link(postID PostID) Link {
	return NewLink(h.Handle, postID)
}

func (h Home) TimelineReadOnly() git.Address {
	return git.NewAddress(h.Handle.URL(), TimelineBranch)
}

func (h Home) TimelineReadWrite() git.Address {
	return git.NewAddress(h.TimelineURL, TimelineBranch)
}

func (h Home) FollowingReadOnly() git.Address {
	return git.NewAddress(h.Handle.URL(), FollowingBranch)
}

func (h Home) FollowingReadWrite() git.Address {
	return git.NewAddress(h.FollowingURL, FollowingBranch)
}
