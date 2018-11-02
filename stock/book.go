package stock

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/covertbert/iex-cli/utils"
	"github.com/dustin/go-humanize"
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
func QueryBook(symbol string, subsection string) {
	if len(symbol) < 1 {
		errors.Error("No argument supplied")
	}

	b := &Book{}
	body := iex.Query("/stock/" + symbol + "/book")

	err := json.Unmarshal(body, &b)

	if err != nil {
		errors.Error("Failed to unmarshal")
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
		{"Market Cap", humanize.Comma(b.Quote.MarketCap)},
		{"PrimaryExchange", b.Quote.PrimaryExchange},
		{"Sector", utils.ReplaceEmptyValue(b.Quote.Sector)},
		{"CalculationPrice", b.Quote.CalculationPrice},
		{"Open", b.Quote.Open},
		{"Open Time", utils.UNIXToHumanReadable(b.Quote.OpenTime)},
		{"Close", b.Quote.Close},
		{"Close Time", utils.UNIXToHumanReadable(b.Quote.CloseTime)},
		{"High", b.Quote.High},
		{"Low", b.Quote.Low},
		{"Latest Price", b.Quote.LatestPrice},
		{"Latest Source", b.Quote.LatestSource},
		{"Latest Time", b.Quote.LatestTime},
		{"Latest Update", utils.UNIXToHumanReadable(b.Quote.LatestUpdate)},
		{"Latest Volume", b.Quote.LatestVolume},
		{"Iex Realtime Price", b.Quote.IexRealtimePrice},
		{"Iex Realtime Size", b.Quote.IexRealtimeSize},
		{"Iex Last Updated", utils.UNIXToHumanReadable(b.Quote.IexLastUpdated)},
		{"Delayed Price", b.Quote.DelayedPrice},
		{"Delayed Price Time", utils.UNIXToHumanReadable(b.Quote.DelayedPriceTime)},
		{"Extended Price", b.Quote.ExtendedPrice},
		{"Extended Change", b.Quote.ExtendedChange},
		{"Extended Change %", b.Quote.ExtendedChangePercent},
		{"Extended Price Time", utils.UNIXToHumanReadable(b.Quote.ExtendedPriceTime)},
		{"Previous Close", b.Quote.PreviousClose},
		{"Change", b.Quote.Change},
		{"Change %", b.Quote.ChangePercent},
		{"Iex Market %", b.Quote.IexMarketPercent},
		{"Iex Volume", b.Quote.IexVolume},
		{"Avg Total Volume", b.Quote.AvgTotalVolume},
		{"Iex Bid Price", b.Quote.IexBidPrice},
		{"Iex Bid Size", b.Quote.IexBidSize},
		{"Iex Ask Price", b.Quote.IexAskPrice},
		{"Iex Ask Size", b.Quote.IexAskSize},
		{"Pe Ratio", b.Quote.PeRatio},
		{"Week 52 High", b.Quote.Week52High},
		{"Week 52 Low", b.Quote.Week52Low},
		{"Ytd Change", b.Quote.YtdChange},
	})
	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}

func bidTable(b Book) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Timestamp", "Price", "Size"})

	for _, bid := range b.Bids {
		t.AppendRows([]table.Row{
			{utils.UNIXToHumanReadable(bid.Timestamp), bid.Price, bid.Size},
		})
	}

	t.SortBy([]table.SortBy{{Name: "Timestamp", Mode: table.Dsc}})

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}

func askTable(b Book) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Timestamp", "Price", "Size"})

	for _, bid := range b.Bids {
		t.AppendRows([]table.Row{
			{utils.UNIXToHumanReadable(bid.Timestamp), bid.Price, bid.Size},
		})
	}

	t.SortBy([]table.SortBy{{Name: "Timestamp", Mode: table.Dsc}})

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}

func tradeTable(b Book) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Timestamp", "TradeID", "Price", "Size", "Is ISO", "Is Odd Lot", "Is Outside Regular Hours", "Is Single Price Cross", "Is Trade Through Exempt"})

	for _, trade := range b.Trades {
		t.AppendRows([]table.Row{
			{utils.UNIXToHumanReadable(trade.Timestamp), trade.TradeID, trade.Price, trade.Size, trade.IsISO, trade.IsOddLot, trade.IsOutsideRegularHours, trade.IsSinglePriceCross, trade.IsTradeThroughExempt},
		})
	}

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	fmt.Println("\r")
	return t
}
