package main

import (
	"flag"
	"io"

	"github.com/b4b4r07/twistd/twistd"
)

const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

var (
	child  = flag.Bool("child", false, "Do run as a daemon")
	config = flag.String("c", "", "Config")
	fg     = flag.Bool("f", false, "Do not run as a daemon")
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) (status int, err error) {
	flag.Parse()

	twistd, err := twistd.NewTwistd(&twistd.Option{
		Child:      *child,
		Config:     *config,
		Foreground: *fg,
	})
	if err != nil {
		return ExitCodeError, err
	}

	// foreground
	if *fg {
		if err := twistd.Run(); err != nil {
			return ExitCodeError, err
		}
		return ExitCodeOK, nil
	}

	if !*child {
		// parent
		if err := twistd.Parent(); err != nil {
			return ExitCodeError, err
		}
	} else {
		// child
		if err := twistd.Child(); err != nil {
			return ExitCodeError, err
		}
	}

	return ExitCodeOK, nil
}
