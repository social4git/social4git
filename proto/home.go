package proto

import "github.com/gov4git/lib4git/git"

type Home struct {
	Handle     Handle
	PublicURL  git.URL
	PrivateURL git.URL
}

func (h Home) Link(postID PostID) Link {
	return NewLink(h.Handle, postID)
}

func (h Home) PublicReadOnly() git.Address {
	return git.NewAddress(h.Handle.GitURL(), PublicBranch)
}

func (h Home) PublicReadWrite() git.Address {
	return git.NewAddress(h.PublicURL, PublicBranch)
}

func (h Home) PrivateReadOnly() git.Address {
	return git.NewAddress(h.PrivateURL, PrivateBranch)
}

func (h Home) PrivateReadWrite() git.Address {
	return git.NewAddress(h.PrivateURL, PrivateBranch)
}
