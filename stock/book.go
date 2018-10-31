package stock

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/covertbert/iex-cli/iex"
)

// Book structure
type Book struct {
	Quote struct {
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
		ExtendedChange        int     `json:"extendedChange"`
		ExtendedChangePercent int     `json:"extendedChangePercent"`
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
	} `json:"quote"`
	Bids []struct {
		Price     float64 `json:"price"`
		Size      int     `json:"size"`
		Timestamp int64   `json:"timestamp"`
	} `json:"bids"`
	Asks []struct {
		Price     float64 `json:"price"`
		Size      int     `json:"size"`
		Timestamp int64   `json:"timestamp"`
	} `json:"asks"`
	Trades []struct {
		Price                 float64 `json:"price"`
		Size                  int     `json:"size"`
		TradeID               int     `json:"tradeId"`
		IsISO                 bool    `json:"isISO"`
		IsOddLot              bool    `json:"isOddLot"`
		IsOutsideRegularHours bool    `json:"isOutsideRegularHours"`
		IsSinglePriceCross    bool    `json:"isSinglePriceCross"`
		IsTradeThroughExempt  bool    `json:"isTradeThroughExempt"`
		Timestamp             int64   `json:"timestamp"`
	} `json:"trades"`
	SystemEvent struct {
		SystemEvent string `json:"systemEvent"`
		Timestamp   int64  `json:"timestamp"`
	} `json:"systemEvent"`
}

// QueryBook returns the pricing infomation for a given company
func QueryBook(ticker string) {
	b := &Book{}
	body := iex.Query("/stock/" + ticker + "/book")

	err := json.Unmarshal(body, &b)

	if err != nil {
		fmt.Println(errors.New("Failed to unmarshal JSON body"))
		return
	}

	fmt.Println(b.Quote.Symbol)
}
