package stock

import (
	"encoding/json"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	ui "github.com/covertbert/termui"
)

// Chart structure
type Chart []struct {
	Date             string  `json:"date"`
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Close            float64 `json:"close"`
	Volume           int64   `json:"volume"`
	UnadjustedVolume int64   `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
	Vwap             float64 `json:"vwap"`
	Label            string  `json:"label"`
	ChangeOverTime   float64 `json:"changeOverTime"`
}

// QueryChart displays live chart information
func QueryChart(symbol string) {
	if noSymbol := len(symbol) < 1; noSymbol {
		errors.Error("No argument supplied")
	}

	c := &Chart{}
	body := iex.Query("/stock/" + symbol + "/chart")

	if err := json.Unmarshal(body, &c); err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	if err := ui.Init(); err != nil {
		errors.Error("Failed to initialize graph", err)
	}

	defer ui.Close()

	lc := ui.NewLineChart()
	lc.BorderLabel = "Close prices"
	lc.BorderFg = ui.ColorWhite
	lc.Data = dataPoints(*c)
	lc.DataLabels = dataLabels(*c)
	lc.Width = 80
	lc.Height = 20
	lc.X = 0
	lc.Y = 0
	lc.AxesColor = ui.ColorWhite
	lc.LineColor["Close"] = ui.ColorBlue
	lc.LineColor["Open"] = ui.ColorYellow
	lc.LineColor["High"] = ui.ColorGreen
	lc.LineColor["Low"] = ui.ColorRed
	lc.Mode = "dot"

	i := ui.NewPar(":PRESS q TO QUIT")
	i.Height = 3
	i.Width = 100
	i.TextFgColor = ui.ColorWhite
	i.BorderLabel = "Keyboard shortcuts"
	i.BorderFg = ui.ColorWhite

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(8, 0, i)),
		ui.NewRow(
			ui.NewCol(8, 0, lc)))

	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}

func dataPoints(c Chart) map[string][]float64 {
	ps := map[string][]float64{}
	ps["Close"] = make([]float64, len(c))
	ps["Open"] = make([]float64, len(c))
	ps["High"] = make([]float64, len(c))
	ps["Low"] = make([]float64, len(c))

	for i, point := range c {
		ps["Close"][i] = point.Close
		ps["Open"][i] = point.Open
		ps["High"][i] = point.High
		ps["Low"][i] = point.Low
	}

	return ps
}

func dataLabels(c Chart) []string {
	p := []string{}

	for _, element := range c {
		p = append(p, element.Date)
	}

	return p
}
