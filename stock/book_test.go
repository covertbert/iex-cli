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

func TestBookQuoteNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "book", "quote")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
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

func TestBookBidsNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "book", "bids")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}

func TestBookAsks(t *testing.T) {
	testcli.Run("../iex-cli", "book", "asks", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Timestamp") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Timestamp")
	}
}

func TestBookAsksNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "book", "asks")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
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

func TestBookTradesNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "book", "trades")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
