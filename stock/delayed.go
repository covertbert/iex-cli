package stock

import (
	"encoding/json"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/jedib0t/go-pretty/table"
)

// Delayed structure
type Delayed struct {
	Symbol           string  `json:"symbol"`
	DelayedPrice     float64 `json:"delayedPrice"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	DelayedSize      int     `json:"delayedSize"`
	DelayedPriceTime int64   `json:"delayedPriceTime"`
	ProcessedTime    int64   `json:"processedTime"`
}

// QueryDelayed displays the 15 minute delayed market quote
func QueryDelayed(symbol string) {
	if len(symbol) < 1 {
		errors.Error("No argument supplied")
	}

	d := &Delayed{}
	body := iex.Query("/stock/" + symbol + "/delayed-quote")
	err := json.Unmarshal(body, &d)

	if err != nil {
		errors.Error("Failed to unmarshal")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Symbol",
		"Delayed Price",
		"High",
		"Low",
		"Delayed Size",
		"Delayed Price Time",
		"Processed Time",
	})
	t.AppendRow(table.Row{
		d.Symbol,
		d.DelayedPrice,
		d.High,
		d.Low,
		d.DelayedSize,
		d.DelayedPriceTime,
		d.ProcessedTime,
	})
	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
