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
		Name:    "book",
		Aliases: []string{"b"},
		Usage:   "view a company's book",
		Action: func(c *cli.Context) error {
			stock.CompanyBook(c.Args().First())
			return nil
		},
	},
	{
		Name:    "company",
		Aliases: []string{"c"},
		Usage:   "view a company's information",
		Action: func(c *cli.Context) error {
			stock.CompanyInfo(c.Args().First())
			return nil
		},
	},
}
