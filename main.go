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
		Usage:   "view a symbol's book",
		Subcommands: []cli.Command{
			{
				Name:    "quote",
				Aliases: []string{"q"},
				Usage:   "view quote",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "quote")
					return nil
				},
			},
			{
				Name:  "bids",
				Usage: "view bids",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "bids")
					return nil
				},
			},
			{
				Name:  "asks",
				Usage: "view asks",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "asks")
					return nil
				},
			},
			{
				Name:  "trades",
				Usage: "view trades",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "trades")
					return nil
				},
			},
		},
	},
	{
		Name:    "company",
		Aliases: []string{"c"},
		Usage:   "view a symbol's general information - e.g. iex-cli company AAPL",
		Action: func(c *cli.Context) error {
			stock.QueryCompany(c.Args().First())
			return nil
		},
	},
	{
		Name:    "ohlc",
		Aliases: []string{"ohlc"},
		Usage:   "view a symbol's official open and close - e.g. iex-cli ohlc AAPL",
		Action: func(c *cli.Context) error {
			stock.QueryOHLC(c.Args().First())
			return nil
		},
	},
}
