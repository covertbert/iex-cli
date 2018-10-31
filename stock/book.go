package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/covertbert/iex-cli/iex"
	"github.com/covertbert/iex-cli/utils"
	"github.com/jedib0t/go-pretty/table"
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
func QueryBook(ticker string, subsection string) {
	b := &Book{}
	body := iex.Query("/stock/" + ticker + "/book")

	err := json.Unmarshal(body, &b)

	if err != nil {
		fmt.Println(errors.New("Failed to unmarshal JSON body"))
		return
	}

	switch subsection {
	case "quote":
		quoteTable(*b).Render()
	case "bids":
		bidTable(*b).Render()
	case "asks":
		askTable(*b).Render()
	case "trades":
		tradeTable(*b).Render()
	}
}

func quoteTable(b Book) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Quote"})
	t.AppendRows([]table.Row{
		{"Symbol", b.Quote.Symbol},
		{"CompanyName", b.Quote.CompanyName},
		{"PrimaryExchange", b.Quote.PrimaryExchange},
		{"Sector", utils.ReplaceEmptyValue(b.Quote.Sector)},
		{"CalculationPrice", b.Quote.CalculationPrice},
		{"Open", b.Quote.Open},
		{"OpenTime", b.Quote.OpenTime},
		{"Close", b.Quote.Close},
		{"CloseTime", b.Quote.CloseTime},
		{"High", b.Quote.High},
		{"Low", b.Quote.Low},
		{"LatestPrice", b.Quote.LatestPrice},
		{"LatestSource", b.Quote.LatestSource},
		{"LatestTime", b.Quote.LatestTime},
		{"LatestUpdate", b.Quote.LatestUpdate},
		{"LatestVolume", b.Quote.LatestVolume},
		{"IexRealtimePrice", b.Quote.IexRealtimePrice},
		{"IexRealtimeSize", b.Quote.IexRealtimeSize},
		{"IexLastUpdated", b.Quote.IexLastUpdated},
		{"DelayedPrice", b.Quote.DelayedPrice},
		{"DelayedPriceTime", b.Quote.DelayedPriceTime},
		{"ExtendedPrice", b.Quote.ExtendedPrice},
		{"ExtendedChange", b.Quote.ExtendedChange},
		{"ExtendedChangePercent", b.Quote.ExtendedChangePercent},
		{"ExtendedPriceTime", b.Quote.ExtendedPriceTime},
		{"PreviousClose", b.Quote.PreviousClose},
		{"Change", b.Quote.Change},
		{"ChangePercent", b.Quote.ChangePercent},
		{"IexMarketPercent", b.Quote.IexMarketPercent},
		{"IexVolume", b.Quote.IexVolume},
		{"AvgTotalVolume", b.Quote.AvgTotalVolume},
		{"IexBidPrice", b.Quote.IexBidPrice},
		{"IexBidSize", b.Quote.IexBidSize},
		{"IexAskPrice", b.Quote.IexAskPrice},
		{"IexAskSize", b.Quote.IexAskSize},
		{"MarketCap", b.Quote.MarketCap},
		{"PeRatio", b.Quote.PeRatio},
		{"Week52High", b.Quote.Week52High},
		{"Week52Low", b.Quote.Week52Low},
		{"YtdChange", b.Quote.YtdChange},
	})
	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}

func bidTable(b Book) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Bids"})

	for i, bid := range b.Bids {
		if i == 0 {
			t.AppendRows([]table.Row{
				{"Timestamp", "Price", "Size"},
				{utils.UNIXToHumanReadable(bid.Timestamp), bid.Price, bid.Size},
			})
		} else {
			t.AppendRows([]table.Row{
				{""},
				{"Timestamp", "Price", "Size"},
				{utils.UNIXToHumanReadable(bid.Timestamp), bid.Price, bid.Size},
			})
		}
	}

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}

func askTable(b Book) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Asks"})

	for i, ask := range b.Asks {
		if i == 0 {
			t.AppendRows([]table.Row{
				{"Timestamp", "Price", "Size"},
				{utils.UNIXToHumanReadable(ask.Timestamp), ask.Price, ask.Size},
			})
		} else {
			t.AppendRows([]table.Row{
				{""},
				{"Timestamp", "Price", "Size"},
				{utils.UNIXToHumanReadable(ask.Timestamp), ask.Price, ask.Size},
			})
		}
	}

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}

func tradeTable(b Book) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Asks"})

	for i, trade := range b.Trades {
		if i == 0 {
			t.AppendRows([]table.Row{
				{"Price", "Size", "TradeID", "IsISO", "IsOddLot", "IsOutsideRegularHours", "IsSinglePriceCross", "IsTradeThroughExempt", "Timestamp"},
				{trade.Price, trade.Size, trade.TradeID, trade.IsISO, trade.IsOddLot, trade.IsOutsideRegularHours, trade.IsSinglePriceCross, trade.IsTradeThroughExempt, utils.UNIXToHumanReadable(trade.Timestamp)},
			})
		} else {
			t.AppendRows([]table.Row{
				{""},
				{"Price", "Size", "TradeID", "IsISO", "IsOddLot", "IsOutsideRegularHours", "IsSinglePriceCross", "IsTradeThroughExempt", "Timestamp"},
				{trade.Price, trade.Size, trade.TradeID, trade.IsISO, trade.IsOddLot, trade.IsOutsideRegularHours, trade.IsSinglePriceCross, trade.IsTradeThroughExempt, utils.UNIXToHumanReadable(trade.Timestamp)},
			})
		}
	}

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}
