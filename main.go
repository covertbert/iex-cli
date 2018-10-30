package main

import (
	"log"
	"os"

	"github.com/covertbert/iex-cli/stock"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "iex-cli"
	app.Usage = "Stock market information in the command line via iex"

	app.Commands = cliCommands

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

var cliCommands = []cli.Command{
	{
		Name:    "company",
		Aliases: []string{"s"},
		Usage:   "query a company",
		Action: func(c *cli.Context) error {
			stock.CompanyInfo(c.Args().First())
			return nil
		},
	},
}
