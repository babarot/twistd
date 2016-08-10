package twistd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Slack struct {
	Text        string        `json:"text"`
	Username    string        `json:"username"`
	Icon_emoji  string        `json:"icon_emoji"`
	Attachments []Attachments `json:"attachments"`
}

type Attachments struct {
	Color      string `json:"color"`
	AuthotName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`
	Title      string `json:"title"`
	TitleLink  string `json:"title_link"`
	Text       string `json:"text"`
	Footer     string `json:"footer"`
	FooterIcon string `json:"footer_icon"`
	TimeStamp  string `json:"ts"`
}

func incomingHooks(slackUrl string, s Slack) error {
	params, _ := json.Marshal(s)
	resp, err := http.PostForm(
		slackUrl,
		url.Values{
			"payload": {string(params)},
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(string(params))

	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (s *Slack) Post() error {
	var conf ConfToml
	if err := LoadConf("/Users/b4b4r07/src/github.com/b4b4r07/twistd/config.toml", &conf); err != nil {
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
		var (
			format = "Mon Jan 02 15:04:05 -0700 2006"
			target = tweet.CreatedAt
		)
		loc, _ := time.LoadLocation("Asia/Tokyo")
		ts, _ := time.ParseInLocation(format, target, loc)

		// Debug Log
		fmt.Println(
			tweet.User.ScreenName,
			"https://twitter.com/"+tweet.User.ScreenName,
			tweet.User.ProfileImageURL,
			tweet.User.Name,
			"https://twitter.com/"+tweet.User.ScreenName+"/status/"+tweet.IDStr,
			tweet.Text,
			fmt.Sprint(ts.Unix()),
		)

		incomingHooks(
			conf.Slack.Url,
			Slack{
				Username:   "Search zplug from Tweets",
				Icon_emoji: ":bird:",
				Attachments: []Attachments{
					Attachments{
						Color:      "#55acee",
						AuthotName: tweet.User.ScreenName,
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
		)
	}

	filterParams := &twitter.StreamFilterParams{
		Track: []string{
			"enhancd",
			"hoge",
			"tomato",
			"word",
			"zplug",
		},
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
	log.Println(<-ch)
	stream.Stop()

	return nil
}
