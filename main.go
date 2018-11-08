package main

import (
	"log"
	"os"

	"github.com/covertbert/iex-cli/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "iex-cli"
	app.Usage = "Stock market information in the command line via iex"
	app.Version = "1.0.0"
	app.Commands = commands.CliCommands

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
