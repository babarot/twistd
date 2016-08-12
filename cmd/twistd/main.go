package main

import (
	"fmt"
	"os"
)

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}

	status, err := cli.Run(os.Args)
	if err != nil {
		fmt.Fprintln(cli.outStream, err)
	}

	os.Exit(status)
}
