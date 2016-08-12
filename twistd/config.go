package twistd

import (
	"github.com/BurntSushi/toml"
)

type ConfToml struct {
	Core    SectionCore    `toml:"core"`
	Slack   SectionSlack   `toml:"slack"`
	Twitter SectionTwitter `toml:"twitter"`
}

type SectionCore struct {
	PidFile    string   `toml:"pid_file"`
	LogFile    string   `toml:"log_file"`
	Words      []string `toml:"words"`
	ForeGround bool     `toml:"fore_ground"`
}

type SectionSlack struct {
	Url       string `toml:"url"`
	Channel   string `toml:"channel"`
	Username  string `toml:"username"`
	IconEmoji string `toml:"icon_emoji"`
}

type SectionTwitter struct {
	ConsumerKey       string `toml:"consumer_key"`
	ConsumerKeySecret string `toml:"consumer_key_secret"`
	AccessToken       string `toml:"access_token"`
	AccessTokenSecret string `toml:"access_token_secret"`
}

func LoadConf(confPath string, confToml *ConfToml) error {
	_, err := toml.DecodeFile(confPath, confToml)
	if err != nil {
		return err
	}
	return nil
}
