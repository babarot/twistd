# twistd

**WIP**

Twitter Streaming Daemon

A daemon to filter with the word you want to search from Tweets via Twitter Streaming API


## Usage

```
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

Options | Description
---|---
-c, --config | Specify config file

Commands | Description
---|---
start | Start to run as a daemon
stop | Stop twistd running as a daemon

## Installation

```
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
pid_file = "/var/log/twistd.log"
log_file = "/var/run/twistd.pid"

[slack]
url = "https://hooks.slack.com/services/..."

[twitter]
consumer_key = "********"
consumer_key_secret = "****************"
access_token = "********"
access_token_secret = "****************"
```
