package markets

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/covertbert/iex-cli/utils"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

// Market structure
type Market []struct {
	Mic           string  `json:"mic"`
	TapeID        string  `json:"tapeId"`
	VenueName     string  `json:"venueName"`
	Volume        int64   `json:"volume"`
	TapeA         int64   `json:"tapeA"`
	TapeB         int64   `json:"tapeB"`
	TapeC         int64   `json:"tapeC"`
	MarketPercent float64 `json:"marketPercent"`
	LastUpdated   int64   `json:"lastUpdated"`
}

// QueryMarkets displays near real time traded volume on the markets
func QueryMarkets() {
	m := &Market{}
	body := iex.Query("/market")
	err := json.Unmarshal(body, &m)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Mic",
		"Tape ID",
		"Venue Name",
		"Volume",
		"Tape A",
		"Tape B",
		"Tape C",
		"Market Percent",
		"Last Updated",
	})

	for _, market := range *m {
		t.AppendRow(table.Row{
			market.Mic,
			utils.ReplaceEmptyValue(market.TapeID),
			market.VenueName,
			humanize.Comma(market.Volume),
			humanize.Comma(market.TapeA),
			humanize.Comma(market.TapeB),
			humanize.Comma(market.TapeC),
			fmt.Sprintf("%.3f", market.MarketPercent),
			utils.UNIXToHumanReadable(market.LastUpdated),
		})
	}

	t.SetAlign([]text.Align{
		text.AlignLeft,
		text.AlignLeft,
		text.AlignLeft,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignLeft,
	})

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
