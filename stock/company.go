package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/covertbert/iex-cli/iex"
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

// CompanyInfo shows general info for a given company
func CompanyInfo(ticker string) {
	c := &Company{}
	body := iex.Query("/stock/" + ticker + "/company")
	err := json.Unmarshal(body, &c)

	if err != nil {
		fmt.Println(errors.New("Failed to unmarshal JSON body"))
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
		replaceEmpty(c.Symbol),
		replaceEmpty(c.CompanyName),
		replaceEmpty(c.Exchange),
		replaceEmpty(c.Industry),
		replaceEmpty(c.Website),
		replaceEmpty(c.CEO),
		replaceEmpty(c.IssueType),
		replaceEmpty(c.Sector),
	})
	t.SetAllowedColumnLengths([]int{40, 40, 40, 40, 40, 40, 40, 40})
	t.SetStyle(table.StyleColoredCyanWhiteOnBlack)
	t.Render()
}

func replaceEmpty(value string) string {
	if len(value) > 0 {
		return value
	}

	return "N/A"
}
