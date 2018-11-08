package stock

import (
	"encoding/json"
	"fmt"
	"strings"

	e "errors"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	ui "github.com/gizak/termui"
)

// Dividends structure
type Dividends []struct {
	ExDate       string  `json:"exDate"`
	PaymentDate  string  `json:"paymentDate"`
	RecordDate   string  `json:"recordDate"`
	DeclaredDate string  `json:"declaredDate"`
	Amount       float64 `json:"amount"`
	Flag         string  `json:"flag"`
	Type         string  `json:"type"`
	Qualified    string  `json:"qualified"`
	Indicated    string  `json:"indicated"`
}

// QueryDividends displays live dividends information
func QueryDividends(symbol string, rng string) {
	if noSymbol := len(symbol) < 1; noSymbol {
		errors.Error("No argument supplied")
	}

	d := &Dividends{}

	path, err := divQueryPath(symbol, rng)

	if err != nil {
		errors.Error("Incorrect arguments passed to dividends command", err)
	}

	body := iex.Query(path)

	if err := json.Unmarshal(body, &d); err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	if noData := len(*d) < 1; noData {
		errors.Error("Currently no data for the specified range")
	}

	if err := ui.Init(); err != nil {
		errors.Error("Failed to initialize graph", err)
	}

	defer ui.Close()

	h := divChartHeader(symbol, rng)
	ld := dividendsData(*d)
	p := dividendsInformation()

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(10, 0, h)),
		ui.NewRow(
			ui.NewCol(10, 0, ld),
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

func divQueryPath(symbol string, rng string) (string, error) {
	allowedRanges := []string{"5y", "2y", "1y", "ytd", "6m", "3m", "1m", "1d", ""}

	for _, allowedRng := range allowedRanges {
		if rng == allowedRng {
			return fmt.Sprintf("/stock/%v/dividends/%v", symbol, rng), nil
		}
	}

	return "", e.New("Range is not valid. Run iex-cli dividends --range for help")
}

func divChartHeader(symbol string, rng string) *ui.Par {
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

func dividendsInformation() *ui.Par {
	i := ui.NewPar(":PRESS q TO QUIT")
	i.Height = 3
	i.Width = 100
	i.TextFgColor = ui.ColorWhite
	i.BorderLabel = "Keyboard shortcuts"
	i.BorderFg = ui.ColorWhite

	return i
}

func dividendsData(d Dividends) *ui.LineChart {
	p := []float64{}

	for _, element := range d {
		p = append(p, element.Amount)
		p = append(p, element.Amount)
		p = append(p, element.Amount)
	}

	lc := ui.NewLineChart()
	lc.BorderLabel = "Dividends"
	lc.BorderFg = ui.ColorWhite
	lc.Data = p
	lc.DataLabels = divDataLabels(d)
	lc.Height = 15
	lc.X = 0
	lc.Y = 0
	lc.AxesColor = ui.ColorWhite
	lc.LineColor = ui.ColorYellow
	lc.Mode = "dot"

	return lc
}

func divDataLabels(d Dividends) []string {
	p := []string{}

	for _, element := range d {
		l := element.ExDate
		p = append(p, l)
		p = append(p, l)
		p = append(p, l)
	}

	return p
}
