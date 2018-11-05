package stock

import (
	"encoding/json"
	"os"

	"github.com/covertbert/iex-cli/utils"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/jedib0t/go-pretty/table"
)

// Sector structure
type Sector []struct {
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Performance float64 `json:"performance"`
	LastUpdated int64   `json:"lastUpdated"`
}

// QuerySector displays a list of each sector and performance for the current trading day
func QuerySector() {
	s := &Sector{}
	body := iex.Query("/stock/market/sector-performance")
	err := json.Unmarshal(body, &s)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Name",
		"Performance",
		"Last updated",
	})

	for _, sector := range *s {
		t.AppendRow(table.Row{
			sector.Name,
			sector.Performance,
			utils.UNIXToHumanReadable(sector.LastUpdated),
		})
	}

	t.SortBy([]table.SortBy{{Name: "Performance", Mode: table.DscNumeric}})

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
