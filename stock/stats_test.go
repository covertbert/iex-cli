package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestStats(t *testing.T) {
	testcli.Run("../iex-cli", "stats", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Symbol") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Symbol")
	}

	if !testcli.StdoutContains("Float") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Float")
	}

	if !testcli.StdoutContains("Debt") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Debt")
	}

	if !testcli.StdoutContains("Beta") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Beta")
	}
}

func TestStatsNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "stats")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
