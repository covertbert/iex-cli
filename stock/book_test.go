package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestBookQuote(t *testing.T) {
	testcli.Run("../iex-cli", "book", "quote", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Symbol") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Symbol")
	}
}
func TestBookBids(t *testing.T) {
	testcli.Run("../iex-cli", "book", "bids", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Timestamp") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Timestamp")
	}
}

func TestBookAsks(t *testing.T) {
	testcli.Run("../iex-cli", "book", "asks", "AAPL")
	// testcli.Run("./iex-cli", "book", "asks", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Timestamp") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Timestamp")
	}
}

func TestBookTrades(t *testing.T) {
	testcli.Run("../iex-cli", "book", "trades", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Timestamp") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Timestamp")
	}
}
