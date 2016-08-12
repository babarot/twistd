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
		//twistd.Logger.Error(
		//	map[string]interface{}{
		//		"message": fmt.Sprint(err),
		//	},
		//)

	}

	os.Exit(status)
}
