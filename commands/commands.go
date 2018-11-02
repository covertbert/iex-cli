package commands

import (
	"github.com/covertbert/iex-cli/stock"
	"github.com/urfave/cli"
)

// CliCommands lists the commands for the CLI app
var CliCommands = []cli.Command{
	{
		Name:    "book",
		Aliases: []string{"b"},
		Usage:   "view a symbol's book - Requires subcommand and symbol e.g. iex-cli book quote AAPL",
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
				Name:    "bids",
				Aliases: []string{"b"},
				Usage:   "view bids",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "bids")
					return nil
				},
			},
			{
				Name:    "asks",
				Aliases: []string{"a"},
				Usage:   "view asks",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "asks")
					return nil
				},
			},
			{
				Name:    "trades",
				Aliases: []string{"t"},
				Usage:   "view trades",
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
		Name:    "crypto",
		Aliases: []string{"cr"},
		Usage:   "view cryptocurrency market information - e.g. iex-cli crypto",
		Action: func(c *cli.Context) error {
			stock.QueryCrypto()
			return nil
		},
	},
	{
		Name:    "delayed",
		Aliases: []string{"d"},
		Usage:   "view a symbol's 15 minute delayed market quote - e.g. iex-cli delayed AAPL",
		Action: func(c *cli.Context) error {
			stock.QueryDelayed(c.Args().First())
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
