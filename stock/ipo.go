package stock

import (
	"encoding/json"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/jedib0t/go-pretty/table"
)

// IPO structure
type IPO struct {
	ViewData []struct {
		Company  string `json:"Company"`
		Symbol   string `json:"Symbol"`
		Price    string `json:"Price"`
		Shares   string `json:"Shares"`
		Amount   string `json:"Amount"`
		Float    string `json:"Float"`
		Percent  string `json:"Percent"`
		Market   string `json:"Market"`
		Expected string `json:"Expected"`
	} `json:"viewData"`
}

// QueryIPO displays upcoming IPO data
func QueryIPO() {
	i := &IPO{}
	body := iex.Query("/stock/market/upcoming-ipos")
	err := json.Unmarshal(body, &i)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Company",
		"Symbol",
		"Price",
		"Shares",
		"Amount",
		"Float",
		"Percent",
		"Market",
		"Expected",
	})

	for _, ipo := range i.ViewData {
		t.AppendRow(table.Row{
			ipo.Company,
			ipo.Symbol,
			ipo.Price,
			ipo.Shares,
			ipo.Amount,
			ipo.Float,
			ipo.Percent,
			ipo.Market,
			ipo.Expected,
		})
	}

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	t.Render()
}
