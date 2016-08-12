package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Slack struct {
	Text        string        `json:"text"`
	Username    string        `json:"username"`
	IconEmoji   string        `json:"icon_emoji"`
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

func Post(slackUrl string, s Slack) error {
	if slackUrl == "" {
		return errors.New("Error: slack incoming url is empty")
	}

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
