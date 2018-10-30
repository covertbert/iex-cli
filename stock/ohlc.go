package stock

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/covertbert/iex-cli/iex"
)

// OHLC structure
type OHLC struct {
	Open struct {
		Price float64 `json:"price"`
		Time  int64   `json:"time"`
	} `json:"open"`
	Close struct {
		Price float64 `json:"price"`
		Time  int64   `json:"time"`
	} `json:"close"`
	High float64 `json:"high"`
	Low  float64 `json:"low"`
}

// QueryOHLC returns the official open and close for a give symbol.
func QueryOHLC(ticker string) {
	o := &OHLC{}
	body := iex.Query("/stock/" + ticker + "/ohlc")
	err := json.Unmarshal(body, &o)

	if err != nil {
		fmt.Println(errors.New("Failed to unmarshal JSON body"))
	}

	fmt.Println(o)
}
