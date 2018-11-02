package iex

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/covertbert/iex-cli/errors"
)

const baseURL = "https://api.iextrading.com/1.0/"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// Query queries the IEX api with the given path
func Query(path string) []byte {
	resp, err := httpClient.Get(baseURL + path)

	if err != nil {
		errors.Error("API request failed")
	}

	if resp.StatusCode != 200 {
		errors.Error("The symbol you have passed does not exist")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		errors.Error("Failed to read response body")
	}

	return body
}
