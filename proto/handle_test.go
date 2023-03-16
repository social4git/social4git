package proto

import "testing"

func TestHandleParse(t *testing.T) {

	for i, raw := range testHandleRaw {
		h, err := ParseHandle(raw)
		if err != nil {
			t.Error(err)
			continue
		}
		if testHandleParsed[i] != h {
			t.Errorf("parse handle %d, expecting %v, got %v", i, testHandleParsed[i].DebugString(), h.DebugString())
		}
	}
}

var testHandleRaw = []string{
	"file:///tmp/petar/1.social4git.public",
	"https://github.com/petar/social4git.public",
}

var testHandleParsed = []Handle{
	{Scheme: "file", Host: "", Path: "/tmp/petar/1.social4git.public"},
	{Scheme: "https", Host: "github.com", Path: "/petar/social4git.public"},
}

func TestHandleString(t *testing.T) {
	for i, h := range testHandleParsed {
		s := h.String()
		if s != testHandleRaw[i] {
			t.Errorf("expecting %v, got %v", testHandleRaw[i], s)
		}
	}
}
