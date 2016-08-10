package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/b4b4r07/twistd/twistd"
)

const (
	DAEMON_START = 1 + iota
	DAEMON_SUCCESS
	DAEMON_FAIL
)

func Parent() error {
	args := []string{"--child"}
	args = append(args, os.Args[1:]...)

	// for connection of child proc
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}

	cmd := exec.Command(os.Args[0], args...)
	cmd.ExtraFiles = []*os.File{w}
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err = cmd.Start(); err != nil {
		return err
	}

	// get child proc info from the pipe
	var status int = DAEMON_START
	go func() {
		buf := make([]byte, 1)
		r.Read(buf)

		if int(buf[0]) == DAEMON_SUCCESS {
			status = int(buf[0])
		} else {
			status = DAEMON_FAIL
		}
	}()

	// run child proc (wait 30s)
	i := 0
	for i < 60 {
		if status != DAEMON_START {
			break
		}
		time.Sleep(500 * time.Millisecond)
		i++
	}

	// terminate parent proc
	if status == DAEMON_SUCCESS {
		return nil
	} else {
		return fmt.Errorf("Child failed to start")
	}
}

func Child() error {
	// write here if any initializations
	var err error

	// notify child proc info to parent proc
	pipe := os.NewFile(uintptr(3), "pipe")
	if pipe != nil {
		defer pipe.Close()
		if err == nil {
			pipe.Write([]byte{DAEMON_SUCCESS})
		} else {
			pipe.Write([]byte{DAEMON_FAIL})
		}
	}

	// ignore SIGCHILD
	signal.Ignore(syscall.SIGCHLD)

	// close each STDOUT, STDIN, STDERR
	syscall.Close(0)
	syscall.Close(1)
	syscall.Close(2)

	// set proc group leader
	syscall.Setsid()

	// claer Umask
	syscall.Umask(022)

	// cd to root
	syscall.Chdir("/")

	// main loop
	var slack twistd.Slack
	if err := slack.Post(); err != nil {
		return err
	}

	return nil
}
