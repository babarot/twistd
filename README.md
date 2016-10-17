# twistd

> *Twitter Streaming Daemon*

![](https://raw.githubusercontent.com/b4b4r07/screenshots/master/twistd/main.png)

A daemon to filter with words you want to search from Tweets via Twitter Streaming API and to send a message to your Slack team

## Usage

```console
$ sudo make install    # copy config to /etc
$ make serve
```

### Interfaces

Option | Description
---|---
`-c` | Specify config file

## Installation

```console
$ git clone https://github.com/b4b4r07/twistd
```

## Configuration

The configuration file (format [`toml`](https://github.com/toml-lang/toml)) is here:

```toml
[core]
words = [
  "tomato",
  "potato",
]
log_file = "/var/log/twistd.log"
pid_file = "/var/run/twistd.pid"

[slack]
channel = "#random"
url = "https://hooks.slack.com/services/..."
username = "Twistd"
icon_emoji = ":bird:"

[twitter]
consumer_key = "********"
consumer_key_secret = "****************"
access_token = "********"
access_token_secret = "****************"
```

## License

[MIT](http://b4b4r07.mit-license.org)

## Author

@b4b4r07
