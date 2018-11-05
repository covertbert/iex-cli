package stock

import (
	"encoding/json"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/covertbert/iex-cli/utils"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
)

// Quote structure
type Quote struct {
	Symbol                string  `json:"symbol"`
	CompanyName           string  `json:"companyName"`
	PrimaryExchange       string  `json:"primaryExchange"`
	Sector                string  `json:"sector"`
	CalculationPrice      string  `json:"calculationPrice"`
	Open                  float64 `json:"open"`
	OpenTime              int64   `json:"openTime"`
	Close                 float64 `json:"close"`
	CloseTime             int64   `json:"closeTime"`
	High                  float64 `json:"high"`
	Low                   float64 `json:"low"`
	LatestPrice           float64 `json:"latestPrice"`
	LatestSource          string  `json:"latestSource"`
	LatestTime            string  `json:"latestTime"`
	LatestUpdate          int64   `json:"latestUpdate"`
	LatestVolume          int     `json:"latestVolume"`
	IexRealtimePrice      float64 `json:"iexRealtimePrice"`
	IexRealtimeSize       int     `json:"iexRealtimeSize"`
	IexLastUpdated        int64   `json:"iexLastUpdated"`
	DelayedPrice          float64 `json:"delayedPrice"`
	DelayedPriceTime      int64   `json:"delayedPriceTime"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	ExtendedPriceTime     int64   `json:"extendedPriceTime"`
	PreviousClose         float64 `json:"previousClose"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	IexMarketPercent      float64 `json:"iexMarketPercent"`
	IexVolume             float64 `json:"iexVolume"`
	AvgTotalVolume        float64 `json:"avgTotalVolume"`
	IexBidPrice           float64 `json:"iexBidPrice"`
	IexBidSize            float64 `json:"iexBidSize"`
	IexAskPrice           float64 `json:"iexAskPrice"`
	IexAskSize            float64 `json:"iexAskSize"`
	MarketCap             int64   `json:"marketCap"`
	PeRatio               float64 `json:"peRatio"`
	Week52High            float64 `json:"week52High"`
	Week52Low             float64 `json:"week52Low"`
	YtdChange             float64 `json:"ytdChange"`
}

// QueryQuote displays general info for a given company by symbol
func QueryQuote(symbol string) {
	if noSymbol := len(symbol) < 1; noSymbol {
		errors.Error("No argument supplied")
	}

	q := &Quote{}
	body := iex.Query("/stock/" + symbol + "/quote")
	err := json.Unmarshal(body, &q)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRows([]table.Row{
		{"Symbol", q.Symbol},
		{"Company Name", q.CompanyName},
		{"Market Cap", humanize.Comma(q.MarketCap)},
		{"Primary Exchange", q.PrimaryExchange},
		{"Sector", utils.ReplaceEmptyValue(q.Sector)},
		{"Calculation Price", q.CalculationPrice},
		{"Open", q.Open},
		{"Open Time", utils.UNIXToHumanReadable(q.OpenTime)},
		{"Close", q.Close},
		{"Close Time", utils.UNIXToHumanReadable(q.CloseTime)},
		{"High", q.High},
		{"Low", q.Low},
		{"Latest Price", q.LatestPrice},
		{"Latest Source", q.LatestSource},
		{"Latest Time", q.LatestTime},
		{"Latest Update", utils.UNIXToHumanReadable(q.LatestUpdate)},
		{"Latest Volume", q.LatestVolume},
		{"Iex Realtime Price", q.IexRealtimePrice},
		{"Iex Realtime Size", q.IexRealtimeSize},
		{"Iex Last Updated", utils.UNIXToHumanReadable(q.IexLastUpdated)},
		{"Delayed Price", q.DelayedPrice},
		{"Delayed Price Time", utils.UNIXToHumanReadable(q.DelayedPriceTime)},
		{"Extended Price", q.ExtendedPrice},
		{"Extended Change", q.ExtendedChange},
		{"Extended Change %", q.ExtendedChangePercent},
		{"Extended Price Time", utils.UNIXToHumanReadable(q.ExtendedPriceTime)},
		{"Previous Close", q.PreviousClose},
		{"Change", q.Change},
		{"Change %", q.ChangePercent},
		{"Iex Market %", q.IexMarketPercent},
		{"Iex Volume", q.IexVolume},
		{"Avg Total Volume", q.AvgTotalVolume},
		{"Iex Bid Price", q.IexBidPrice},
		{"Iex Bid Size", q.IexBidSize},
		{"Iex Ask Price", q.IexAskPrice},
		{"Iex Ask Size", q.IexAskSize},
		{"Pe Ratio", q.PeRatio},
		{"Week 52 High", q.Week52High},
		{"Week 52 Low", q.Week52Low},
		{"Ytd Change", q.YtdChange},
	})

	t.SortBy([]table.SortBy{{Number: 1, Mode: table.Asc}})

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
