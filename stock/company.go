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

// QueryCompany shows general info for a given company by symbol
func QueryCompany(ticker string) {
	c := &Company{}
	body := iex.Query("/stock/" + ticker + "/company")
	err := json.Unmarshal(body, &c)

	if err != nil {
		fmt.Println(errors.New("Failed to unmarshal JSON body"))
		return
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
		utils.ReplaceEmpty(c.Symbol),
		utils.ReplaceEmpty(c.CompanyName),
		utils.ReplaceEmpty(c.Exchange),
		utils.ReplaceEmpty(c.Industry),
		utils.ReplaceEmpty(c.Website),
		utils.ReplaceEmpty(c.CEO),
		utils.ReplaceEmpty(c.IssueType),
		utils.ReplaceEmpty(c.Sector),
	})
	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)

	t.Render()
}
