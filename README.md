# twistd

> *Twitter Streaming Daemon*

![](https://raw.githubusercontent.com/b4b4r07/screenshots/master/twistd/main.png)

A daemon to filter with words you want to search from Tweets via Twitter Streaming API and to send a message to your Slack team

## Usage

```console
$ /etc/init.d/twistd start
```

### Interfaces

For more infomation, see also [`run.sh`](run.sh) and [`Makefile`](Makefile).

Interface | Description
---|---
start |
stop |
restart |
status |

There are twistd's interfaces.

Option | Description
---|---
-c | Specify config file
-f | Do not run as a daemon

## Installation

```console
$ git clone https://github.com/b4b4r07/twistd /home/you
$ cd /home/you/twistd
$ sudo make install
```

## Configuration

The configuration file (format [`toml`](https://github.com/toml-lang/toml)) is here:

```toml
[core]
fore_ground = false
words = [
  "red",
  "yellow",
  "green",
]
log_file = "/var/log/twistd.log"
pid_file = "/var/run/twistd.pid"

[slack]
url = "https://hooks.slack.com/services/..."
username = "Twistd"
icon_emoji = ":bird:"

[twitter]
consumer_key = "********"
consumer_key_secret = "****************"
access_token = "********"
access_token_secret = "****************"
```

## TODOs

- [x] Logging system
- [ ] Pid Lock
- [ ] Improve `run.sh` (init script)
