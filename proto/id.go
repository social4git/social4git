package proto

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gov4git/lib4git/ns"
)

type PostID struct {
	Time        time.Time
	ContentHash string
	Nonce       string
}

func NewPostID(t time.Time, content []byte) PostID {
	return PostID{Time: t.UTC(), ContentHash: ContentHash(content), Nonce: ContentHash(Nonce())}
}

const IDTimeFormat = "20060102150405"

func ParsePostID(s string) (PostID, error) {
	ps := strings.Split(s, "_")
	if len(ps) != 3 {
		return PostID{}, fmt.Errorf("unexpected number of parts in post id")
	}
	t, err := time.Parse(IDTimeFormat, ps[0])
	if err != nil {
		return PostID{}, err
	}
	return PostID{
		Time:        t,
		ContentHash: ps[1],
		Nonce:       ps[2],
	}, nil
}

func (x PostID) String() string {
	t := x.Time.Format(IDTimeFormat)
	return t + "_" + x.ContentHash + "_" + x.Nonce
}

func (x PostID) NS() ns.NS {
	return append(PostDayNS(x.Time), x.String())
}

func PostDayNS(t time.Time) ns.NS {
	year := fmt.Sprintf("%04d", t.Year())
	month := fmt.Sprintf("%02d", t.Month())
	day := fmt.Sprintf("%02d", t.Day())
	return ns.NS{PostDir, year, month, day}
}

type PostIDs []PostID

func (ps PostIDs) Sort() {
	sort.Sort(ps)
}

func (ps PostIDs) Len() int {
	return len(ps)
}

func (ps PostIDs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps PostIDs) Less(i, j int) bool {
	return ps[i].String() < ps[j].String()
}

func MapToPostIDs(m map[PostID]bool) PostIDs {
	r := PostIDs{}
	for id := range m {
		r = append(r, id)
	}
	return r
}
