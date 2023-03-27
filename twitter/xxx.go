package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	twitter "github.com/g8rswimmer/go-twitter/v2"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func main() {
	token := flag.String("token", "", "twitter API token")
	flag.Parse()

	client := &twitter.Client{
		Authorizer: authorize{
			Token: *token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	ctx := context.Background()
	opts := twitter.UserTweetTimelineOpts{
		Expansions: []twitter.Expansion{
			twitter.ExpansionEntitiesMentionsUserName,
			twitter.ExpansionAuthorID,
			twitter.ExpansionGeoPlaceID,
			twitter.ExpansionInReplyToUserID,
		},
		MediaFields: []twitter.MediaField{
			twitter.MediaFieldURL,
		},
		TweetFields: []twitter.TweetField{
			twitter.TweetFieldID,
			twitter.TweetFieldText,
			twitter.TweetFieldAuthorID,
			twitter.TweetFieldContextAnnotations,
			twitter.TweetFieldConversationID,
			twitter.TweetFieldCreatedAt,
			twitter.TweetFieldGeo,
			twitter.TweetFieldInReplyToUserID,
			twitter.TweetFieldLanguage,
			twitter.TweetFieldAttachments,
		},
		MaxResults: 10e3,
		StartTime:  time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:    time.Now(),
		// PaginationToken: XXX,
		// SinceID:         XXX,
		// UntilID:         XXX,
	}
	timeline, err := client.UserTweetTimeline(ctx, "maymounkov", opts)
	if err != nil {
		log.Panic(err)
	}

	enc, err := json.MarshalIndent(timeline, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}
