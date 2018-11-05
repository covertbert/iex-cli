package refdata

import (
	"encoding/json"
	"os"

	"github.com/covertbert/iex-cli/errors"

	"github.com/covertbert/iex-cli/iex"
	"github.com/jedib0t/go-pretty/table"
)

// Symbol structure
type Symbol []struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	IsEnabled bool   `json:"isEnabled"`
	Type      string `json:"type"`
}

type flexInt [][]interface{}

// QuerySymbols displays a list of all symbols on IEX
func QuerySymbols() {
	s := &Symbol{}
	body := iex.Query("/ref-data/symbols")
	err := json.Unmarshal(body, &s)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Symbol",
		"Name",
		"Date",
		"Is Enabled",
		"Type",
	})

	for _, symbol := range *s {
		t.AppendRow(table.Row{
			symbol.Symbol,
			symbol.Name,
			symbol.Date,
			symbol.IsEnabled,
			symbol.Type,
		})
	}

	t.SetAllowedColumnLengths([]int{40, 80, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	t.Render()
}
