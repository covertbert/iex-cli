package stock

import (
	"encoding/json"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/table"
)

// Key structure
type Key struct {
	CompanyName         string  `json:"companyName"`
	Marketcap           int64   `json:"marketcap"`
	Beta                float64 `json:"beta"`
	Week52High          float64 `json:"week52high"`
	Week52Low           float64 `json:"week52low"`
	Week52Change        float64 `json:"week52change"`
	ShortInterest       int     `json:"shortInterest"`
	ShortDate           int     `json:"shortDate"`
	DividendRate        float64 `json:"dividendRate"`
	DividendYield       float64 `json:"dividendYield"`
	ExDividendDate      string  `json:"exDividendDate"`
	LatestEPS           int     `json:"latestEPS"`
	LatestEPSDate       string  `json:"latestEPSDate"`
	SharesOutstanding   int64   `json:"sharesOutstanding"`
	Float               int64   `json:"float"`
	ReturnOnEquity      float64 `json:"returnOnEquity"`
	ConsensusEPS        float64 `json:"consensusEPS"`
	NumberOfEstimates   int     `json:"numberOfEstimates"`
	EPSSurpriseDollar   float64 `json:"EPSSurpriseDollar"`
	EPSSurprisePercent  float64 `json:"EPSSurprisePercent"`
	Symbol              string  `json:"symbol"`
	EBITDA              int64   `json:"EBITDA"`
	Revenue             int64   `json:"revenue"`
	GrossProfit         int64   `json:"grossProfit"`
	Cash                int64   `json:"cash"`
	Debt                int64   `json:"debt"`
	TtmEPS              float64 `json:"ttmEPS"`
	RevenuePerShare     int     `json:"revenuePerShare"`
	RevenuePerEmployee  int     `json:"revenuePerEmployee"`
	PeRatioHigh         float64 `json:"peRatioHigh"`
	PeRatioLow          int     `json:"peRatioLow"`
	ReturnOnAssets      float64 `json:"returnOnAssets"`
	ReturnOnCapital     float64 `json:"returnOnCapital"`
	ProfitMargin        float64 `json:"profitMargin"`
	PriceToSales        float64 `json:"priceToSales"`
	PriceToBook         float64 `json:"priceToBook"`
	Day200MovingAvg     float64 `json:"day200MovingAvg"`
	Day50MovingAvg      float64 `json:"day50MovingAvg"`
	InstitutionPercent  float64 `json:"institutionPercent"`
	InsiderPercent      float64 `json:"insiderPercent"`
	ShortRatio          float64 `json:"shortRatio"`
	Year5ChangePercent  float64 `json:"year5ChangePercent"`
	Year2ChangePercent  float64 `json:"year2ChangePercent"`
	Year1ChangePercent  float64 `json:"year1ChangePercent"`
	YtdChangePercent    float64 `json:"ytdChangePercent"`
	Month6ChangePercent float64 `json:"month6ChangePercent"`
	Month3ChangePercent float64 `json:"month3ChangePercent"`
	Month1ChangePercent float64 `json:"month1ChangePercent"`
	Day5ChangePercent   float64 `json:"day5ChangePercent"`
	Day30ChangePercent  float64 `json:"day30ChangePercent"`
}

// QueryKey displays key information for a symbol
func QueryKey(symbol string) {
	if noSymbol := len(symbol) < 1; noSymbol {
		errors.Error("No argument supplied")
	}

	k := &Key{}
	body := iex.Query("/stock/" + symbol + "/stats")
	err := json.Unmarshal(body, &k)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRows([]table.Row{
		{"Symbol", k.Symbol},
		{"Company Name", k.CompanyName},
		{"Marketcap", humanize.Comma(k.Marketcap)},
		{"Beta", k.Beta},
		{"Week 52 High", k.Week52High},
		{"Week 52 Low", k.Week52Low},
		{"Week 52 Change", k.Week52Change},
		{"Short Interest", k.ShortInterest},
		{"Short Date", k.ShortDate},
		{"Dividend Rate", k.DividendRate},
		{"Dividend Yield", k.DividendYield},
		{"Ex Dividend Date", k.ExDividendDate},
		{"Latest EPS", k.LatestEPS},
		{"Latest EPS Date", k.LatestEPSDate},
		{"Shares Outstanding", k.SharesOutstanding},
		{"Float", k.Float},
		{"Return On Equity", k.ReturnOnEquity},
		{"Consensus EPS", k.ConsensusEPS},
		{"Number Of Estimates", k.NumberOfEstimates},
		{"EPS Surprise Dollar", k.EPSSurpriseDollar},
		{"EPS Surprise Percent", k.EPSSurprisePercent},
		{"EBITDA", humanize.Comma(k.EBITDA)},
		{"Revenue", humanize.Comma(k.Revenue)},
		{"Gross Profit", humanize.Comma(k.GrossProfit)},
		{"Cash", humanize.Comma(k.Cash)},
		{"Debt", humanize.Comma(k.Debt)},
		{"Ttm EPS", k.TtmEPS},
		{"Revenue Per Share", k.RevenuePerShare},
		{"Revenue Per Employee", k.RevenuePerEmployee},
		{"Pe Ratio High", k.PeRatioHigh},
		{"Pe Ratio Low", k.PeRatioLow},
		{"Return On Assets", k.ReturnOnAssets},
		{"Return On Capital", k.ReturnOnCapital},
		{"Profit Margin", k.ProfitMargin},
		{"Price To Sales", k.PriceToSales},
		{"Price To Book", k.PriceToBook},
		{"Day 200 Moving Avg", k.Day200MovingAvg},
		{"Day 50 Moving Avg", k.Day50MovingAvg},
		{"Institution Percent", k.InstitutionPercent},
		{"Insider Percent", k.InsiderPercent},
		{"Short Ratio", k.ShortRatio},
		{"Year 5 Change Percent", k.Year5ChangePercent},
		{"Year 2 Change Percent", k.Year2ChangePercent},
		{"Year 1 ChangePercent", k.Year1ChangePercent},
		{"Ytd Change Percent", k.YtdChangePercent},
		{"Month 6 Change Percent", k.Month6ChangePercent},
		{"Month 3 Change Percent", k.Month3ChangePercent},
		{"Month 1 Change Percent", k.Month1ChangePercent},
		{"Day 5 Change Percent", k.Day5ChangePercent},
		{"Day 30 Change Percent", k.Day30ChangePercent},
	})

	t.SortBy([]table.SortBy{{Number: 1, Mode: table.Asc}})

	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
