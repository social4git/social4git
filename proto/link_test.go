package proto

import (
	"testing"
	"time"
)

func TestLinkParse(t *testing.T) {
	for i, raw := range testLinkRaw {
		l, err := ParseLink(raw)
		if err != nil {
			t.Error(err)
			continue
		}
		if l != testLinkParsed[i] {
			t.Errorf("expecting %v, got %v", testLinkParsed[i], l)
		}
	}
}

var testLinkRaw = []string{
	"social4git-file:///tmp/petar/1.social4git.public?post=20230316170915_33112ee14ee469c3eb52fe90322ec81dd404a0093d565a6d71ce77cbc8124e3b_1afcd13ab72aca9ae7d9af3ec1e3082cc82de991dc69274a926f23a29d7f27ef",
	"social4git-https://github.com/petar/social4git.public?post=20230316170915_33112ee14ee469c3eb52fe90322ec81dd404a0093d565a6d71ce77cbc8124e3b_1afcd13ab72aca9ae7d9af3ec1e3082cc82de991dc69274a926f23a29d7f27ef",
}

var testLinkParsed = []Link{
	{
		Handle: Handle{Scheme: "file", Host: "", Path: "/tmp/petar/1.social4git.public"},
		PostID: PostID{
			Time:        time.Date(2023, 03, 16, 17, 9, 15, 0, time.UTC),
			ContentHash: "33112ee14ee469c3eb52fe90322ec81dd404a0093d565a6d71ce77cbc8124e3b",
			Nonce:       "1afcd13ab72aca9ae7d9af3ec1e3082cc82de991dc69274a926f23a29d7f27ef",
		},
	},
	{
		Handle: Handle{Scheme: "https", Host: "github.com", Path: "/petar/social4git.public"},
		PostID: PostID{
			Time:        time.Date(2023, 03, 16, 17, 9, 15, 0, time.UTC),
			ContentHash: "33112ee14ee469c3eb52fe90322ec81dd404a0093d565a6d71ce77cbc8124e3b",
			Nonce:       "1afcd13ab72aca9ae7d9af3ec1e3082cc82de991dc69274a926f23a29d7f27ef",
		},
	},
}

func TestLinkString(t *testing.T) {
	for i, l := range testLinkParsed {
		s := l.String()
		if s != testLinkRaw[i] {
			t.Errorf("expecting %v, got %v", testLinkRaw[i], s)
		}
	}
}
