package commands

import (
	"github.com/covertbert/iex-cli/markets"
	"github.com/covertbert/iex-cli/refdata"
	"github.com/covertbert/iex-cli/stock"
	"github.com/urfave/cli"
)

// CliCommands lists the commands for the CLI app
var CliCommands = []cli.Command{
	{
		Name:    "book",
		Aliases: []string{"b"},
		Usage:   "View a symbol's book",
		Subcommands: []cli.Command{
			{
				Name:    "quote",
				Aliases: []string{"q"},
				Usage:   "View quote",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "quote")
					return nil
				},
			},
			{
				Name:    "bids",
				Aliases: []string{"b"},
				Usage:   "View bids",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "bids")
					return nil
				},
			},
			{
				Name:    "asks",
				Aliases: []string{"a"},
				Usage:   "View asks",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "asks")
					return nil
				},
			},
			{
				Name:    "trades",
				Aliases: []string{"t"},
				Usage:   "View trades",
				Action: func(c *cli.Context) error {
					stock.QueryBook(c.Args().First(), "trades")
					return nil
				},
			},
		},
	},
	{
		Name:    "chart",
		Aliases: []string{"c"},
		Usage:   "View a symbol's chart",
		Action: func(c *cli.Context) error {
			stock.QueryChart(c.Args().First(), c.String("range"))
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "range",
				Usage: "Chart range - options: 5y, 2y, 1y, ytd, 6m, 3m, 1m, 1d, date/<UNIX Timestamp>, dynamic",
			},
		},
	},
	{
		Name:    "company",
		Aliases: []string{"c"},
		Usage:   "View a symbol's general information",
		Action: func(c *cli.Context) error {
			stock.QueryCompany(c.Args().First())
			return nil
		},
	},
	{
		Name:    "crypto",
		Aliases: []string{"cr"},
		Usage:   "View cryptocurrency market information",
		Action: func(c *cli.Context) error {
			stock.QueryCrypto()
			return nil
		},
	},
	{
		Name:    "delayed",
		Aliases: []string{"d"},
		Usage:   "View a symbol's 15 minute delayed market quote",
		Action: func(c *cli.Context) error {
			stock.QueryDelayed(c.Args().First())
			return nil
		},
	},
	{
		Name:    "ipo",
		Aliases: []string{"i"},
		Usage:   "View upcoming IPO information",
		Action: func(c *cli.Context) error {
			stock.QueryIPO()
			return nil
		},
	},
	{
		Name:    "markets",
		Aliases: []string{"m"},
		Usage:   "View a list of markets with volume",
		Action: func(c *cli.Context) error {
			markets.QueryMarkets()
			return nil
		},
	},
	{
		Name:    "news",
		Aliases: []string{"n"},
		Usage:   "View stock market news",
		Action: func(c *cli.Context) error {
			stock.QueryNews(c.Args().First())
			return nil
		},
	},
	{
		Name:    "ohlc",
		Aliases: []string{"ohlc"},
		Usage:   "View a symbol's official open and close",
		Action: func(c *cli.Context) error {
			stock.QueryOHLC(c.Args().First())
			return nil
		},
	},
	{
		Name:    "peers",
		Aliases: []string{"pe"},
		Usage:   "View a list of peer tickers",
		Action: func(c *cli.Context) error {
			stock.QueryPeers(c.Args().First())
			return nil
		},
	},
	{
		Name:    "price",
		Aliases: []string{"pr"},
		Usage:   "View a symbol's current price",
		Action: func(c *cli.Context) error {
			stock.QueryPrice(c.Args().First())
			return nil
		},
	},
	{
		Name:    "quote",
		Aliases: []string{"q"},
		Usage:   "View a symbol's quote information",
		Action: func(c *cli.Context) error {
			stock.QueryQuote(c.Args().First())
			return nil
		},
	},
	{
		Name:    "sector",
		Aliases: []string{"se"},
		Usage:   "View each sector's performance for the current trading day",
		Action: func(c *cli.Context) error {
			stock.QuerySector()
			return nil
		},
	},
	{
		Name:    "stats",
		Aliases: []string{"st"},
		Usage:   "View a symbol's key stats",
		Action: func(c *cli.Context) error {
			stock.QueryKey(c.Args().First())
			return nil
		},
	},
	{
		Name:    "symbols",
		Aliases: []string{"sy"},
		Usage:   "View a list of symbols",
		Action: func(c *cli.Context) error {
			refdata.QuerySymbols()
			return nil
		},
	},
}
