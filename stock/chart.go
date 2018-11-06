package stock

import (
	"encoding/json"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	ui "github.com/gizak/termui"
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

	lc := lineChart(*c)
	p := chartInformation()

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(8, 0, p)),
		ui.NewRow(
			ui.NewCol(8, 0, lc)))

	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}

func lineChart(c Chart) *ui.LineChart {
	lc := ui.NewLineChart()
	lc.BorderLabel = "Close prices"
	lc.BorderFg = ui.ColorWhite
	lc.Data = dataPoints(c)
	lc.DataLabels = dataLabels(c)
	lc.Width = 100
	lc.Height = 20
	lc.X = 0
	lc.Y = 0
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorBlue
	lc.Mode = "dot"

	return lc
}

func chartInformation() *ui.Par {
	i := ui.NewPar(":PRESS q TO QUIT")
	i.Height = 3
	i.Width = 100
	i.TextFgColor = ui.ColorWhite
	i.BorderLabel = "Keyboard shortcuts"
	i.BorderFg = ui.ColorWhite

	return i
}

func dataPoints(c Chart) []float64 {
	p := []float64{}

	for _, element := range c {
		p = append(p, element.Close)
		p = append(p, element.Close)
		p = append(p, element.Close)
	}

	return p
}

func dataLabels(c Chart) []string {
	p := []string{}

	for _, element := range c {
		p = append(p, element.Date)
		p = append(p, element.Date)
		p = append(p, element.Date)
	}

	return p
}
