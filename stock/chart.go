package stock

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	e "errors"

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
func QueryChart(symbol string, rng string) {
	if noSymbol := len(symbol) < 1; noSymbol {
		errors.Error("No argument supplied")
	}

	c := &Chart{}

	if noData := len(*c) < 1; noData {
		errors.Error("Currently no data for the specified rangeg")
	}

	path, err := queryPath(symbol, rng)

	if err != nil {
		errors.Error("Incorrect arguments passed to chart command", err)
	}

	body := iex.Query(path)

	if err := json.Unmarshal(body, &c); err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	if err := ui.Init(); err != nil {
		errors.Error("Failed to initialize graph", err)
	}

	defer ui.Close()

	h := chartHeader(symbol, rng)
	lo := openData(*c)
	lc := closeData(*c)
	lh := highData(*c)
	ll := lowData(*c)
	p := chartInformation()

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(10, 0, h)),
		ui.NewRow(
			ui.NewCol(5, 0, lo),
			ui.NewCol(5, 0, lc),
		),
		ui.NewRow(
			ui.NewCol(5, 0, lh),
			ui.NewCol(5, 0, ll),
		),
		ui.NewRow(ui.NewCol(10, 0, p)),
	)

	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Loop()
}

func queryPath(symbol string, rng string) (string, error) {
	allowedRanges := []string{"5y", "2y", "1y", "ytd", "6m", "3m", "1m", "1d", ""}

	for _, allowedRng := range allowedRanges {
		if rng == allowedRng {
			return fmt.Sprintf("/stock/%v/chart/%v", symbol, rng), nil
		}
	}

	if len(rng) == 8 {
		_, err := strconv.ParseInt(rng, 10, 64)
		if err != nil {
			return "", e.New("Range is not a valid 8 digit date. Must be in format yyyymmdd e.g. 20180129")
		}
		return fmt.Sprintf("/stock/%v/chart/date/%v", symbol, rng), nil
	}

	return "", e.New("Range is not valid. Run iex-cli chart --range for help")
}

func chartHeader(symbol string, rng string) *ui.Par {
	if noRangeSpecified := len(rng) == 0; noRangeSpecified {
		rng = "default"
	}

	headerString := fmt.Sprintf("Range: %v", rng)

	i := ui.NewPar(fmt.Sprintf("%v", headerString))
	i.Height = 3
	i.Width = 100
	i.TextFgColor = ui.ColorWhite
	i.BorderLabel = strings.ToUpper(symbol)
	i.BorderFg = ui.ColorWhite

	return i
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

func openData(c Chart) *ui.LineChart {
	p := []float64{}

	for _, element := range c {
		p = append(p, element.Open)
		p = append(p, element.Open)
	}

	lc := ui.NewLineChart()
	lc.BorderLabel = "Open"
	lc.BorderFg = ui.ColorWhite
	lc.Data = p
	lc.DataLabels = dataLabels(c)
	lc.Height = 15
	lc.X = 0
	lc.Y = 0
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorYellow
	lc.Mode = "dot"

	return lc
}

func closeData(c Chart) *ui.LineChart {
	p := []float64{}

	for _, element := range c {
		p = append(p, element.Close)
		p = append(p, element.Close)
	}

	lc := ui.NewLineChart()
	lc.BorderLabel = "Close"
	lc.BorderFg = ui.ColorWhite
	lc.Data = p
	lc.DataLabels = dataLabels(c)
	lc.Height = 15
	lc.X = 0
	lc.Y = 0
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorBlue
	lc.Mode = "dot"

	return lc
}

func highData(c Chart) *ui.LineChart {
	p := []float64{}

	for _, element := range c {
		p = append(p, element.High)
		p = append(p, element.High)
	}

	lc := ui.NewLineChart()
	lc.BorderLabel = "High"
	lc.BorderFg = ui.ColorWhite
	lc.Data = p
	lc.DataLabels = dataLabels(c)
	lc.Height = 15
	lc.X = 0
	lc.Y = 0
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorGreen
	lc.Mode = "dot"

	return lc
}

func lowData(c Chart) *ui.LineChart {
	p := []float64{}

	for _, element := range c {
		p = append(p, element.Low)
		p = append(p, element.Low)
	}

	lc := ui.NewLineChart()
	lc.BorderLabel = "Low"
	lc.BorderFg = ui.ColorWhite
	lc.Data = p
	lc.DataLabels = dataLabels(c)
	lc.Height = 15
	lc.X = 0
	lc.Y = 0
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorRed
	lc.Mode = "dot"

	return lc
}

func dataLabels(c Chart) []string {
	p := []string{}

	for _, element := range c {
		d := element.Date
		p = append(p, d)
		p = append(p, d)
	}

	return p
}
