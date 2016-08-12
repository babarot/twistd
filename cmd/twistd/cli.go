package main

import (
	//"errors"
	"flag"
	//"fmt"
	"io"

	"github.com/b4b4r07/twistd/daemon"
	//"github.com/b4b4r07/twistd/twistd"
)

const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) (status int, err error) {
	var child *bool = flag.Bool("child", false, "Run as a child process")
	//var confPath *string = flag.String("config", "", "configuration file")
	flag.Parse()

	/*
		//map[string]interface{}
		twistd.Logger.Error(
			map[string]interface{}{
				"child": *child,
				//"message": fmt.Sprint(err),
				"message": "error",
			},
		)
	*/

	//if flag.NArg() == 0 {
	//	return ExitCodeError, errors.New("To run as a daemon, run by giving the start command.")
	//}

	//switch flag.Arg(0) {
	//case "start":
	if !*child {
		// parent
		if err := daemon.Parent(); err != nil {
			return ExitCodeError, err
		}
	} else {
		// child
		if err := daemon.Child(); err != nil {
			return ExitCodeError, err
		}
	}
	//case "stop":
	//default:
	//}

	return ExitCodeOK, nil
}
