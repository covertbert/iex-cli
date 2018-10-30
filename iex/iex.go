package iex

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://api.iextrading.com/1.0/"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// QueryCompany queries info for a given company
func QueryCompany(ticker string) []byte {
	path := "/stock/" + ticker + "/company"

	resp, err := httpClient.Get(baseURL + path)

	if err != nil {
		fmt.Println(errors.New("Failed to query company by ticker"))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(errors.New("Failed to read response body for company query"))
	}

	return body
}
