package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestQuote(t *testing.T) {
	testcli.Run("../iex-cli", "quote", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Symbol") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Symbol")
	}
}

func TestQuoteNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "quote")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
