package stock

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/covertbert/iex-cli/utils"

	"github.com/covertbert/iex-cli/errors"

	"github.com/covertbert/iex-cli/iex"
)

// News structure
type News []struct {
	Datetime string `json:"datetime"`
	Headline string `json:"headline"`
	Source   string `json:"source"`
	URL      string `json:"url"`
	Summary  string `json:"summary"`
	Related  string `json:"related"`
	Image    string `json:"image"`
}

// QueryNews displays stock market news
func QueryNews(symbol string) {
	var body []byte
	n := &News{}

	body = iex.Query("stock/market/news/last/5")

	if hasSymbol := len(symbol) > 1; hasSymbol {
		body = iex.Query("stock/" + symbol + "/news/last/5")
	}

	err := json.Unmarshal(body, &n)

	if err != nil {
		errors.Error("Failed to unmarshal")
	}

	for _, news := range *n {
		fmt.Println(formatDate(news.Datetime))
		fmt.Println(formatHeadline(news.Headline))
		fmt.Println(formatSummary(news.Summary))
	}
}

func formatDate(date string) string {
	c := color.New(color.FgHiGreen).Add(color.Bold).SprintfFunc()
	return fmt.Sprintf("\n%v", c(utils.DateStringToHumanReadable(date)))
}

func formatHeadline(headline string) string {
	c := color.New(color.FgWhite).Add(color.Bold).SprintfFunc()
	return fmt.Sprintf("\n%v", c(strings.TrimSpace(headline)))
}

func formatSummary(summary string) string {
	c := color.New(color.FgWhite).SprintfFunc()
	return fmt.Sprintf("\n%v\n\n\n", c(strings.TrimSpace(summary)))
}
