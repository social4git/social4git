package cmd

import (
	"fmt"
	"time"

	"github.com/gov4git/lib4git/must"
	"github.com/social4git/social4git/proto"
	"github.com/spf13/cobra"
)

var (
	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show posts",
		Long: `
		Show displays posts from a timeline.

		If the option --my is used, posts published by this user are displayed.
		Otherwise, posts of the users followed by this user are displayed.

		The command needs two additional pieces of information: a window duration (day, month, year) and
		a date inside the window. The command shows posts within the calendar UTC day/month/year
		containing the given date.

		The duration is specified using --day, --month, or --year, defaulting to --month.
		The date is specified with --date, defaulting to today.
`,
		Run: func(cmd *cobra.Command, args []string) {
			h := setup.Home
			fetched := []proto.PostWithMeta{}
			if showMy {
				if showDate != "" {
					date, err := time.Parse(DateLayout, showDate)
					must.NoError(ctx, err)
					if showDay {
						fetched = proto.FetchPublishedPostsDay(ctx, h, date)
					} else if showMonth {
						fetched = proto.FetchPublishedPostsMonth(ctx, h, date)
					} else if showYear {
						fetched = proto.FetchPublishedPostsYear(ctx, h, date)
					} else {
						fetched = proto.FetchPublishedPostsMonth(ctx, h, date)
					}
				} else {
					if showDay {
						fetched = proto.FetchPublishedLatestPostsDay(ctx, h)
					} else if showMonth {
						fetched = proto.FetchPublishedLatestPostsMonth(ctx, h)
					} else if showYear {
						fetched = proto.FetchPublishedLatestPostsYear(ctx, h)
					} else {
						fetched = proto.FetchPublishedLatestPostsMonth(ctx, h)
					}
				}
			} else {
				if showDate != "" {
					date, err := time.Parse(DateLayout, showDate)
					must.NoError(ctx, err)
					if showDay {
						fetched = proto.FetchFollowedPostsDay(ctx, h, date)
					} else if showMonth {
						fetched = proto.FetchFollowedPostsMonth(ctx, h, date)
					} else if showYear {
						fetched = proto.FetchFollowedPostsYear(ctx, h, date)
					} else {
						fetched = proto.FetchFollowedPostsMonth(ctx, h, date)
					}
				} else {
					if showDay {
						fetched = proto.FetchFollowedLatestPostsDay(ctx, h)
					} else if showMonth {
						fetched = proto.FetchFollowedLatestPostsMonth(ctx, h)
					} else if showYear {
						fetched = proto.FetchFollowedLatestPostsYear(ctx, h)
					} else {
						fetched = proto.FetchFollowedLatestPostsMonth(ctx, h)
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
	showCmd.Flags().BoolVar(&showMonth, "month", false, "show a month of posts")
	showCmd.Flags().BoolVar(&showYear, "year", false, "show a year of posts")
	showCmd.Flags().StringVar(&showDate, "date", "", "show posts from a UTC date in format MM/DD/YYYY")
}
