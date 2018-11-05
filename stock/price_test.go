package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestPrice(t *testing.T) {
	testcli.Run("../iex-cli", "price", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
}

func TestPriceNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "price")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
