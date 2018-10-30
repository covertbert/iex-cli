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
		Usage:   "view a ticker's book - e.g. iex-cli book AAPL",
		Action: func(c *cli.Context) error {
			stock.QueryBook(c.Args().First())
			return nil
		},
	},
	{
		Name:    "company",
		Aliases: []string{"c"},
		Usage:   "view a ticker's general information - e.g. iex-cli company AAPL",
		Action: func(c *cli.Context) error {
			stock.QueryCompany(c.Args().First())
			return nil
		},
	},
	{
		Name:    "ohlc",
		Aliases: []string{"ohlc"},
		Usage:   "view a ticker's official open and close - e.g. iex-cli ohlc AAPL",
		Action: func(c *cli.Context) error {
			stock.QueryBook(c.Args().First())
			return nil
		},
	},
}
