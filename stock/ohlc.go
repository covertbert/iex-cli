package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/covertbert/iex-cli/utils"

	"github.com/covertbert/iex-cli/iex"
	"github.com/jedib0t/go-pretty/table"
)

// OHLC structure
type OHLC struct {
	Open struct {
		Price float64 `json:"price"`
		Time  int64   `json:"time"`
	} `json:"open"`
	Close struct {
		Price float64 `json:"price"`
		Time  int64   `json:"time"`
	} `json:"close"`
	High float64 `json:"high"`
	Low  float64 `json:"low"`
}

// QueryOHLC returns the official open and close for a give symbol.
func QueryOHLC(ticker string) {
	o := &OHLC{}
	body := iex.Query("/stock/" + ticker + "/ohlc")
	err := json.Unmarshal(body, &o)

	if err != nil {
		fmt.Println(errors.New("Failed to unmarshal JSON body"))
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Key",
		"Open",
		"Close",
		"High",
		"Low",
	})
	t.AppendRow(table.Row{
		"Price",
		o.Open.Price,
		o.Close.Price,
		o.High,
		o.Low,
	})
	t.AppendRow(table.Row{
		"Time",
		utils.UNIXToHumanReadable(o.Open.Time),
		utils.UNIXToHumanReadable(o.Close.Time),
		"N/A",
		"N/A",
	})
	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	t.Render()
}