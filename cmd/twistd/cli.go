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
	config = flag.String("c", "", "Config")
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) (status int, err error) {
	flag.Parse()

	twistd, err := twistd.NewTwistd(&twistd.Option{
		Config: *config,
	})
	if err != nil {
		return ExitCodeError, err
	}

	if err := twistd.Run(); err != nil {
		return ExitCodeError, err
	}

	return ExitCodeOK, nil
}
