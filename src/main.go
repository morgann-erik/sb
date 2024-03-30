package main

import (
	"log"
	"os"

	"github.com/morgann-erik/sb/setup"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{Commands: []*cli.Command{
        {Name: "setup", Usage: "", Description: "", Action: setup.SetupCommand},
	}}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
