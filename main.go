package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"fmt"
)

const Version = "0.1"

func main() {
	app := cli.NewApp()
	app.Name = "csvcut"
	app.Usage = "cut(1) for comma(*) seperated values"
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag {

	}

	app.Run(os.Args)
}

func run(ctx *cli.Context) error {
	fmt.Printf("Hello world\n")
	return nil
}
