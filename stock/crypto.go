package stock

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/covertbert/iex-cli/utils"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

// Crypto quote structure
type Crypto []struct {
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
	LatestVolume          float64 `json:"latestVolume"`
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
	IexVolume             int     `json:"iexVolume"`
	AvgTotalVolume        int     `json:"avgTotalVolume"`
	IexBidPrice           float64 `json:"iexBidPrice"`
	IexBidSize            int     `json:"iexBidSize"`
	IexAskPrice           float64 `json:"iexAskPrice"`
	IexAskSize            int     `json:"iexAskSize"`
	MarketCap             int64   `json:"marketCap"`
	PeRatio               float64 `json:"peRatio"`
	Week52High            float64 `json:"week52High"`
	Week52Low             float64 `json:"week52Low"`
	YtdChange             float64 `json:"ytdChange"`
	BidPrice              float64 `json:"bidPrice"`
	BidSize               float64 `json:"bidSize"`
	AskPrice              float64 `json:"askPrice"`
	AskSize               float64 `json:"askSize"`
}

// QueryCrypto returns quotes for all Cryptocurrencies supported by IEX
func QueryCrypto() {
	c := &Crypto{}
	body := iex.Query("/stock/market/crypto")
	err := json.Unmarshal(body, &c)

	if err != nil {
		errors.Error(err, "Failed to unmarshal JSON body")
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Symbol",
		"High",
		"Low",
		"Latest",
		"Change",
		"Change%",
		"Bid price",
		"Bid size",
		"Ask price",
		"Ask size",
	})

	for _, crypto := range *c {
		t.AppendRow(table.Row{
			utils.ReplaceEmptyValue(crypto.Symbol),
			fmt.Sprintf("%.3f", crypto.High),
			fmt.Sprintf("%.3f", crypto.Low),
			fmt.Sprintf("%.3f", crypto.LatestPrice),
			fmt.Sprintf("%.3f", crypto.Change),
			fmt.Sprintf("%.3f", crypto.ChangePercent),
			fmt.Sprintf("%.3f", crypto.BidPrice),
			fmt.Sprintf("%.1f", crypto.BidSize),
			fmt.Sprintf("%.3f", crypto.AskPrice),
			fmt.Sprintf("%.1f", crypto.AskSize),
		})
	}

	t.SortBy([]table.SortBy{{Name: "High", Mode: table.DscNumeric}})

	t.SetAlign([]text.Align{
		text.AlignLeft,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
		text.AlignRight,
	})

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
