package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestOHLC(t *testing.T) {
	testcli.Run("../iex-cli", "ohlc", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Key") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Key")
	}

	if !testcli.StdoutContains("Open") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Open")
	}

	if !testcli.StdoutContains("Close") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Close")
	}

	if !testcli.StdoutContains("High") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "High")
	}

	if !testcli.StdoutContains("Low") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Low")
	}
}

func TestOHLCNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "ohlc")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
