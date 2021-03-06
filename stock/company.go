package stock

import (
	"encoding/json"
	"os"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
	"github.com/covertbert/iex-cli/utils"
	"github.com/jedib0t/go-pretty/table"
)

// Company structure
type Company struct {
	Symbol      string   `json:"symbol"`
	CompanyName string   `json:"companyName"`
	Exchange    string   `json:"exchange"`
	Industry    string   `json:"industry"`
	Website     string   `json:"website"`
	Description string   `json:"description"`
	CEO         string   `json:"CEO"`
	IssueType   string   `json:"issueType"`
	Sector      string   `json:"sector"`
	Tags        []string `json:"tags"`
}

// QueryCompany displays general info for a given company by symbol
func QueryCompany(symbol string) {
	if noSymbol := len(symbol) < 1; noSymbol {
		errors.Error("No argument supplied")
	}

	c := &Company{}
	body := iex.Query("/stock/" + symbol + "/company")
	err := json.Unmarshal(body, &c)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Symbol",
		"Company name",
		"Exchange",
		"Industry",
		"Website",
		"CEO",
		"Issue type",
		"Sector",
	})
	t.AppendRow(table.Row{
		utils.ReplaceEmptyValue(c.Symbol),
		utils.ReplaceEmptyValue(c.CompanyName),
		utils.ReplaceEmptyValue(c.Exchange),
		utils.ReplaceEmptyValue(c.Industry),
		utils.ReplaceEmptyValue(c.Website),
		utils.ReplaceEmptyValue(c.CEO),
		utils.ReplaceEmptyValue(c.IssueType),
		utils.ReplaceEmptyValue(c.Sector),
	})
	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
