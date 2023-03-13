package cmd

import (
	"fmt"
	"time"

	"github.com/gov4git/lib4git/must"
	"github.com/petar/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show posts",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			h := setup.Home
			fetched := []proto.PostWithMeta{}
			if showMy {
				if showDate != "" {
					date, err := time.Parse(DateLayout, showDate)
					must.NoError(ctx, err)
					if showDay {
						fetched = proto.FetchTimelinePostsDay(ctx, h, date)
					} else if showMonth {
						fetched = proto.FetchTimelinePostsMonth(ctx, h, date)
					} else {
						fetched = proto.FetchTimelinePostsYear(ctx, h, date)
					}
				} else {
					if showDay {
						fetched = proto.FetchTimelineLatestPostsDay(ctx, h)
					} else if showMonth {
						fetched = proto.FetchTimelineLatestPostsMonth(ctx, h)
					} else {
						fetched = proto.FetchTimelineLatestPostsYear(ctx, h)
					}
				}
			} else {
				if showDate != "" {
					date, err := time.Parse(DateLayout, showDate)
					must.NoError(ctx, err)
					if showDay {
						fetched = proto.FetchFollowingPostsDay(ctx, h, date)
					} else if showMonth {
						fetched = proto.FetchFollowingPostsMonth(ctx, h, date)
					} else {
						fetched = proto.FetchFollowingPostsYear(ctx, h, date)
					}
				} else {
					if showDay {
						fetched = proto.FetchFollowingLatestPostsDay(ctx, h)
					} else if showMonth {
						fetched = proto.FetchFollowingLatestPostsMonth(ctx, h)
					} else {
						fetched = proto.FetchFollowingLatestPostsYear(ctx, h)
					}
				}
			}
			for _, pm := range fetched {
				fmt.Println(pm)
			}
		},
	}
)

var (
	showMy    bool
	showDay   bool
	showMonth bool
	showYear  bool
	showDate  string
)

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().BoolVar(&showMy, "my", false, "if set show my posts, otherwise show posts of users I follow")
	showCmd.Flags().BoolVar(&showDay, "day", false, "show a day of posts")
	showCmd.Flags().BoolVar(&showMonth, "month", true, "show a month of posts")
	showCmd.Flags().BoolVar(&showYear, "year", false, "show a year of posts")
	showCmd.Flags().StringVar(&showDate, "date", "", "show posts from a date")
}
