package twistd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/b4b4r07/twistd/slack"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func (twistd *Twistd) Run() error {
	var conf ConfToml
	if err := LoadConf(twistd.Option.Config, &conf); err != nil {
		return err
	}

	var (
		config = oauth1.NewConfig(
			conf.Twitter.ConsumerKey,
			conf.Twitter.ConsumerKeySecret,
		)
		token = oauth1.NewToken(
			conf.Twitter.AccessToken,
			conf.Twitter.AccessTokenSecret,
		)
	)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if conf.Twitter.SkipRetweet {
			if tweet.RetweetedStatus != nil {
				return
			}
		}
		for _, user := range conf.Twitter.IgnoreUsers {
			switch user.(type) {
			case string:
				if tweet.User.ScreenName == user {
					return
				}
			case int64:
				if tweet.User.ID == user {
					return
				}
			}
		}

		var (
			format = "Mon Jan 02 15:04:05 -0700 2006"
			target = tweet.CreatedAt
		)
		loc, _ := time.LoadLocation("Asia/Tokyo")
		ts, _ := time.ParseInLocation(format, target, loc)

		if err := slack.Post(
			conf.Slack.Url,
			slack.Slack{
				Channel:   conf.Slack.Channel,
				Username:  conf.Slack.Username,
				IconEmoji: conf.Slack.IconEmoji,
				Attachments: []slack.Attachments{
					slack.Attachments{
						Color:      "#55acee",
						AuthotName: "@" + tweet.User.ScreenName,
						AuthorLink: "https://twitter.com/" + tweet.User.ScreenName,
						AuthorIcon: tweet.User.ProfileImageURL,
						Title:      tweet.User.Name + "'s tweet!",
						TitleLink:  "https://twitter.com/" + tweet.User.ScreenName + "/status/" + tweet.IDStr,
						Text:       tweet.Text,
						Footer:     "Twitter",
						FooterIcon: "http://www.freeiconspng.com/uploads/twitter-icon-download-18.png",
						TimeStamp:  fmt.Sprint(ts.Unix()),
					},
				},
			},
		); err != nil {
			twistd.Error(
				map[string]interface{}{
					"message": fmt.Sprint(err),
				},
			)
		}
	}

	filterParams := &twitter.StreamFilterParams{
		Track:         conf.Core.Words,
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		return err
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-ch)
	stream.Stop()

	return nil
}
