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

// Query queries the IEX api with the given path
func Query(path string) []byte {
	resp, err := httpClient.Get(baseURL + path)

	if err != nil {
		fmt.Println(errors.New("HTTP request failed"))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(errors.New("Failed to read response body"))
	}

	return body
}
