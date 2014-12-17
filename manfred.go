package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "manfredctl"
	app.Usage = "Control the Manfred Daemon"
	app.Action = func(c *cli.Context) {
		print("Hello friend!")
	}
	app.Run(os.Args)
}
