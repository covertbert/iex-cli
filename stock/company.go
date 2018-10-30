package stock

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/covertbert/iex-cli/iex"
)

// Company structure
type Company struct {
	Symbol      string
	CompanyName string
	Exchange    string
	Industry    string
	Website     string
	Description string
	CEO         string
	IssueType   string
	Sector      string
	Tags        []string
}

// CompanyInfo renders info for a given company
func CompanyInfo(ticker string) {
	var company Company

	body := iex.QueryCompany(ticker)

	err := json.Unmarshal(body, &company)

	if err != nil {
		fmt.Println(errors.New("Failed to unmarshal JSON body"))
	}

	fmt.Println(company)
}
