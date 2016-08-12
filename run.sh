#!/bin/sh
### BEGIN INIT INFO
# Provides:          twistd
# Required-Start:    $local_fs $network $named $time $syslog
# Required-Stop:     $local_fs $network $named $time $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Description:       Twitter Streaming Daemon
### END INIT INFO

SCRIPT="twistd"
RUNAS="root"
# e.g. CMD="/Users/b4b4r07/src/github.com/b4b4r07/twistd/cmd/twistd/twistd -c $HOME/config.toml"
CMD="$SCRIPT"

PIDFILE=/var/run/${SCRIPT}.pid
LOGFILE=/var/log/${SCRIPT}.log

start() {
    if [ -f /var/run/$PIDNAME ] && kill -0 $(cat /var/run/$PIDNAME) &>/dev/nul; then
        echo "$SCRIPT is already running" >&2
        return 1
    fi
    su - "$RUNAS" -c "$CMD" #& pid=$!
    #echo "$pid">$PIDFILE
    if [ $? -eq 0 ]; then
        echo -e "Starting $SCRIPT:              [\033[32m  OK  \033[m]"
    else
        echo -e "Starting $SCRIPT:              [\033[31m FAIL \033[m]"
    fi
    ps x | grep "twistd" | grep -v "grep" | awk '{print $1}' >"$PIDFILE"
}

stop() {
    if [ ! -f "$PIDFILE" ] || ! kill -0 $(cat "$PIDFILE") &>/dev/null; then
        echo "$SCRIPT is not running" >&2
        return 1
    fi
    kill -15 $(cat "$PIDFILE") && rm -f "$PIDFILE"
    if [ $? -eq 0 ]; then
        echo -e "Starting $SCRIPT:              [\033[32m  OK  \033[m]"
    else
        echo -e "Starting $SCRIPT:              [\033[31m FAIL \033[m]"
    fi
}

case "$1" in
    start)
        start
        exit $?
        ;;
    stop)
        stop
        exit $?
        ;;
    restart)
        stop && start
        exit $?
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|}"
        exit 1
esac
