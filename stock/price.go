package stock

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"

	"github.com/covertbert/iex-cli/errors"
	"github.com/covertbert/iex-cli/iex"
)

// Price structure
type Price float64

// QueryPrice displays a symbol's current price
func QueryPrice(symbol string) {
	if noSymbol := len(symbol) < 1; noSymbol {
		errors.Error("No argument supplied")
	}

	var p Price
	body := iex.Query("/stock/" + symbol + "/price")
	err := json.Unmarshal(body, &p)

	if err != nil {
		errors.Error("Failed to unmarshal", err)
	}

	greenBold := color.New(color.FgHiGreen).Add(color.Bold).SprintFunc()
	price := fmt.Sprintf("\n%v", greenBold(p))

	fmt.Println(price)
}
